package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

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

	topic := "demo"
	for i := 0; i < 5; i++ {
		msg := "message number " + string(i+'0')
		_, err := client.Publish(context.Background(), &pb.Message{Topic: topic, Data: msg})
		if err != nil {
			log.Println("failed to publish: ", err)
		}
		log.Println("Published:", msg)
		time.Sleep(2 * time.Second)
	}
}
