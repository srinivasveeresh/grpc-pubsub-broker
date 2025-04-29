package main

import (
	"context"
	"crypto/tls"
	"log"
	"math"
	"os"
	"strings"
	"time"

	pb "github.com/srinivasveeresh/grpc-pubsub-broker/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func connectClient(addresses []string, useTLS bool) pb.PubSubClient {
	baseDelay := time.Second
	maxDelay := 30 * time.Second
	for {
		for _, addr := range addresses {
			var opts []grpc.DialOption
			if useTLS {
				creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
				opts = append(opts, grpc.WithTransportCredentials(creds))
			} else {
				opts = append(opts, grpc.WithInsecure())
			}
			conn, err := grpc.Dial(addr, opts...)
			if err == nil {
				return pb.NewPubSubClient(conn)
			}
			log.Printf("Failed to connect to %s, will retry: %v", addr, err)
		}
		for attempt := 1; attempt < 5; attempt++ {
			delay := time.Duration(math.Pow(2, float64(attempt))) * baseDelay
			if delay > maxDelay {
				delay = maxDelay
			}
			log.Printf("Backing off for %v before retrying...", delay)
			time.Sleep(delay)
		}
	}
}

func main() {
	addresses := strings.Split(os.Getenv("SERVERS"), ",")
	useTLS := os.Getenv("TLS") == "true"
	client := connectClient(addresses, useTLS)

	ctx := context.Background()

	stream, err := client.Subscribe(ctx, &pb.SubscribeRequest{Topic: "foo"})
	if err != nil {
		log.Fatalf("failed to subscribe: %v", err)
	}

	initial, _ := stream.Recv()
	log.Println("Subscribed with ID:", initial.SubscriberId)

	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Received:", msg.Data)
		}
	}()

	time.Sleep(2 * time.Second)

	client.Publish(ctx, &pb.Message{Topic: "foo", Data: "Hello from publisher!"})

	time.Sleep(10 * time.Second)
}
