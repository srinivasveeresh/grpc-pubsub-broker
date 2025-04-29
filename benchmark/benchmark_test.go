package benchmark

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"

	pb "github.com/srinivasveeresh/grpc-pubsub-broker/proto"
	"google.golang.org/grpc"
)

func BenchmarkHighVolumePublish(b *testing.B) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		b.Skip("Server not running locally; skipping benchmark")
	}
	defer conn.Close()
	client := pb.NewPubSubClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.Subscribe(ctx, &pb.SubscribeRequest{Topic: "benchmark"})
	if err != nil {
		b.Fatalf("failed to subscribe: %v", err)
	}
	_, _ = stream.Recv()

	var wg sync.WaitGroup
	numMessages := 10000
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < numMessages; i++ {
			_, err := client.Publish(context.Background(), &pb.Message{Topic: "benchmark", Data: "ping"})
			if err != nil {
				log.Printf("publish failed: %v", err)
			}
		}
	}()

	received := 0
	for received < numMessages {
		_, err := stream.Recv()
		if err != nil {
			b.Fatalf("stream receive failed: %v", err)
		}
		received++
	}

	wg.Wait()
}
