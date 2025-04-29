package main

import (
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/srinivasveeresh/grpc-pubsub-broker/broker"
	pb "github.com/srinivasveeresh/grpc-pubsub-broker/proto"
	"google.golang.org/grpc"
)

func startTestServer(t *testing.T, address string) *grpc.Server {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPubSubServer(s, &server{br: broker.New()})
	go func() {
		s.Serve(lis)
	}()
	return s
}

func TestServerPublishSubscribe(t *testing.T) {
	address := ":7001"
	s := startTestServer(t, address)
	defer s.Stop()

	time.Sleep(500 * time.Millisecond)

	conn, err := grpc.Dial("localhost"+address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()
	client := pb.NewPubSubClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := client.Subscribe(ctx, &pb.SubscribeRequest{Topic: "test"})
	if err != nil {
		t.Fatalf("subscribe failed: %v", err)
	}
	_, _ = stream.Recv()

	_, err = client.Publish(context.Background(), &pb.Message{Topic: "test", Data: "hello test"})
	if err != nil {
		t.Fatalf("publish failed: %v", err)
	}

	recv, err := stream.Recv()
	if err != nil {
		t.Fatalf("recv failed: %v", err)
	}
	if recv.Data != "hello test" {
		t.Fatalf("unexpected data: %v", recv.Data)
	}
}

func TestServerPublishAndMetrics(t *testing.T) {
	address := ":7101"
	s := startTestServer(t, address)
	defer s.Stop()

	time.Sleep(500 * time.Millisecond)

	conn, err := grpc.Dial("localhost"+address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewPubSubClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := client.Subscribe(ctx, &pb.SubscribeRequest{Topic: "testmetrics"})
	if err != nil {
		t.Fatalf("failed to subscribe: %v", err)
	}
	_, _ = stream.Recv()

	_, err = client.Publish(ctx, &pb.Message{Topic: "testmetrics", Data: "metric test 1"})
	if err != nil {
		t.Fatalf("failed to publish: %v", err)
	}

	recv, err := stream.Recv()
	if err != nil {
		t.Fatalf("failed to receive: %v", err)
	}
	if recv.Data != "metric test 1" {
		t.Fatalf("unexpected data: %v", recv.Data)
	}

	time.Sleep(500 * time.Millisecond)

	resp, err := http.Get("http://localhost:2112/metrics")
	if err != nil {
		t.Fatalf("failed to GET metrics: %v", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	metricsText := string(body)

	if !strings.Contains(metricsText, `pubsub_messages_published_total{topic="testmetrics"} 1`) {
		t.Fatalf("metrics output missing expected publish count")
	}
	if !strings.Contains(metricsText, `pubsub_subscribers_connected 1`) {
		t.Fatalf("metrics output missing expected subscriber count")
	}
}
