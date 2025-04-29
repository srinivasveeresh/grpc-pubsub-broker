package main

import (
	"context"
	"log"
	"os"
	"strings"

	pb "github.com/srinivasveeresh/grpc-pubsub-broker/proto"
	"google.golang.org/grpc"
)

func main() {
	addresses := strings.Split(os.Getenv("SERVERS"), ",")
	conn, err := grpc.Dial(addresses[0], grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to connect: ", err)
	}
	defer conn.Close()

	client := pb.NewPubSubClient(conn)

	stream, err := client.Subscribe(context.Background(), &pb.SubscribeRequest{Topic: "demo"})
	if err != nil {
		log.Fatal("failed to subscribe: ", err)
	}

	initial, _ := stream.Recv()
	log.Println("Subscribed with ID:", initial.SubscriberId)

	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Received:", msg.Data)
	}
}
