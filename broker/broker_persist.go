package broker

import (
	"encoding/json"
	"fmt"
	"os"

	pb "github.com/srinivasveeresh/grpc-pubsub-broker/proto"
)

func ensureDataFolder() {
	_ = os.MkdirAll("data", 0755)
}

func (b *Broker) persistMessage(topic, msg string) {
	ensureDataFolder()
	f, err := os.OpenFile(fmt.Sprintf("data/messages_%s.log", topic), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("%s\n", msg))
}

func (b *Broker) persistSubscribers(meta map[string]*pb.SubscriberMetadata) {
	ensureDataFolder()
	f, err := os.OpenFile("data/subscribers.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.Encode(meta)
}

func (b *Broker) loadSubscribers() map[string]*pb.SubscriberMetadata {
	file, err := os.Open("data/subscribers.json")
	if err != nil {
		return make(map[string]*pb.SubscriberMetadata)
	}
	defer file.Close()

	var meta map[string]*pb.SubscriberMetadata
	err = json.NewDecoder(file).Decode(&meta)
	if err != nil {
		return make(map[string]*pb.SubscriberMetadata)
	}
	return meta
}

// // Extend server main.go startup to reload subscribers

// func (s *server) reloadSubscribersOnStart() {
// 	s.subscriberMeta = s.br.loadSubscribers()
// }

// // Add inside server main()
// 	myServer := &server{br: broker.New()}
// 	pb.RegisterPubSubServer(s, myServer)
// 	myServer.reloadSubscribersOnStart()

// ---
