package broker

import (
	"sync"

	"github.com/google/uuid"
)

type Broker struct {
	Subscribers map[string]map[string]chan string // topic -> subscriberID -> channel
	mu          sync.Mutex
}

type Subscription struct {
	ID     string
	Chan   <-chan string
	Cancel func()
}

func New() *Broker {
	return &Broker{
		Subscribers: make(map[string]map[string]chan string),
	}
}

func (b *Broker) Subscribe(topic string) Subscription {
	c := make(chan string, 10)
	id := uuid.New().String()
	b.mu.Lock()
	if b.Subscribers[topic] == nil {
		b.Subscribers[topic] = make(map[string]chan string)
	}
	b.Subscribers[topic][id] = c
	b.mu.Unlock()

	cancelFunc := func() {
		b.mu.Lock()
		defer b.mu.Unlock()
		delete(b.Subscribers[topic], id)
		close(c)
	}

	return Subscription{ID: id, Chan: c, Cancel: cancelFunc}
}

func (b *Broker) Publish(topic, msg string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	for _, ch := range b.Subscribers[topic] {
		select {
		case ch <- msg:
		default:
		}
	}
	go b.persistMessage(topic, msg)
	go b.trimOldMessages(topic)
}

func (b *Broker) PublishToSubscriber(topic, subscriberID, msg string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if sub, exists := b.Subscribers[topic][subscriberID]; exists {
		select {
		case sub <- msg:
		default:
		}
	}
}
