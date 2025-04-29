package main

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/srinivasveeresh/grpc-pubsub-broker/broker"
	pb "github.com/srinivasveeresh/grpc-pubsub-broker/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
)

type server struct {
	pb.UnimplementedPubSubServer
	br             *broker.Broker
	peerConns      []pb.PubSubClient
	mu             sync.Mutex
	subscriberMeta map[string]*pb.SubscriberMetadata
}

func (s *server) Publish(ctx context.Context, msg *pb.Message) (*pb.Ack, error) {
	s.br.Publish(msg.Topic, msg.Data)
	for _, peer := range s.peerConns {
		_, _ = peer.Forward(context.Background(), msg)
	}
	return &pb.Ack{Status: "ok"}, nil
}

func (s *server) PublishToSubscriber(ctx context.Context, req *pb.TargetedMessage) (*pb.Ack, error) {
	s.br.PublishToSubscriber(req.Topic, req.SubscriberId, req.Data)
	return &pb.Ack{Status: "ok"}, nil
}

func (s *server) PublishToSubscribers(ctx context.Context, req *pb.TargetedGroupMessage) (*pb.Ack, error) {
	for _, id := range req.SubscriberIds {
		s.br.PublishToSubscriber(req.Topic, id, req.Data)
	}
	return &pb.Ack{Status: "ok"}, nil
}

func (s *server) Forward(ctx context.Context, msg *pb.Message) (*pb.Ack, error) {
	s.br.Publish(msg.Topic, msg.Data)
	return &pb.Ack{Status: "ok"}, nil
}

func (s *server) Subscribe(req *pb.SubscribeRequest, stream pb.PubSub_SubscribeServer) error {
	sub := s.br.Subscribe(req.Topic)
	defer sub.Cancel()

	p, ok := peer.FromContext(stream.Context())
	ip := "unknown"
	if ok {
		ip = p.Addr.String()
	}

	s.mu.Lock()
	if s.subscriberMeta == nil {
		s.subscriberMeta = make(map[string]*pb.SubscriberMetadata)
	}
	s.subscriberMeta[sub.ID] = &pb.SubscriberMetadata{
		SubscriberId: sub.ID,
		IpAddress:    ip,
		Tags:         []string{},
	}
	s.mu.Unlock()

	stream.Send(&pb.SubscriptionMessage{SubscriberId: sub.ID, Topic: req.Topic, Data: ""})

	heartbeatTimeout := 60 * time.Second
	ticker := time.NewTicker(heartbeatTimeout)
	defer ticker.Stop()

	for {
		select {
		case msg, ok := <-sub.Chan:
			if !ok {
				return nil
			}
			if err := stream.Send(&pb.SubscriptionMessage{SubscriberId: sub.ID, Topic: req.Topic, Data: msg}); err != nil {
				return err
			}
		case <-stream.Context().Done():
			return nil
		case <-ticker.C:
			return nil
		}
	}
}

func (s *server) ListSubscribers(ctx context.Context, req *pb.SubscribeRequest) (*pb.SubscriberList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	list := make([]*pb.SubscriberMetadata, 0)
	if subs, ok := s.br.Subscribers[req.Topic]; ok {
		for id := range subs {
			meta := s.subscriberMeta[id]
			list = append(list, meta)
		}
	}
	return &pb.SubscriberList{Subscribers: list}, nil
}

func (s *server) ListTopics(ctx context.Context, _ *pb.Empty) (*pb.TopicList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	topics := make([]string, 0, len(s.br.Subscribers))
	for topic := range s.br.Subscribers {
		topics = append(topics, topic)
	}
	return &pb.TopicList{Topics: topics}, nil
}

func (s *server) Unsubscribe(ctx context.Context, req *pb.UnsubscribeRequest) (*pb.Ack, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if subs, ok := s.br.Subscribers[req.Topic]; ok {
		if ch, found := subs[req.SubscriberId]; found {
			close(ch)
			delete(subs, req.SubscriberId)
			delete(s.subscriberMeta, req.SubscriberId)
		}
	}
	return &pb.Ack{Status: "ok"}, nil
}

func main() {
	address := os.Getenv("ADDRESS")
	peers := strings.Split(os.Getenv("PEERS"), ",")
	useTLS := os.Getenv("TLS") == "true"

	var opts []grpc.ServerOption
	if useTLS {
		creds, err := credentials.NewServerTLSFromFile("cert.pem", "key.pem")
		if err != nil {
			log.Fatalf("Failed to load TLS keys: %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(opts...)
	myServer := &server{br: broker.New()}
	pb.RegisterPubSubServer(s, myServer)

	for _, peer := range peers {
		if peer == "" {
			continue
		}
		go func(peerAddr string) {
			for {
				var dialOpts []grpc.DialOption
				if useTLS {
					creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
					dialOpts = append(dialOpts, grpc.WithTransportCredentials(creds))
				} else {
					dialOpts = append(dialOpts, grpc.WithInsecure())
				}
				conn, err := grpc.Dial(peerAddr, dialOpts...)
				if err == nil {
					myServer.peerConns = append(myServer.peerConns, pb.NewPubSubClient(conn))
					log.Printf("Connected to peer: %s", peerAddr)
					return
				}
				log.Printf("Retrying peer connection to %s", peerAddr)
				time.Sleep(3 * time.Second)
			}
		}(peer)
	}

	log.Printf("Server listening at %v", address)
	s.Serve(lis)
}
