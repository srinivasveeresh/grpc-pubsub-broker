package main

import (
	"context"
	"testing"
	"time"

	pb "github.com/srinivasveeresh/grpc-pubsub-broker/proto"
	"google.golang.org/grpc"
)

func TestClientConnectSubscribePublish(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Skip("Server not running locally; skipping client test")
	}
	defer conn.Close()
	client := pb.NewPubSubClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := client.Subscribe(ctx, &pb.SubscribeRequest{Topic: "clienttest"})
	if err != nil {
		t.Fatalf("failed to subscribe: %v", err)
	}
	_, _ = stream.Recv()

	_, err = client.Publish(ctx, &pb.Message{Topic: "clienttest", Data: "hello from client"})
	if err != nil {
		t.Fatalf("failed to publish: %v", err)
	}

	recv, err := stream.Recv()
	if err != nil {
		t.Fatalf("failed to receive: %v", err)
	}
	if recv.Data != "hello from client" {
		t.Fatalf("unexpected data received: %v", recv.Data)
	}
}
