package broker

import (
	"testing"
	"time"
)

func TestBrokerSubscribeAndPublish(t *testing.T) {
	b := New()
	topic := "test_topic"
	sub := b.Subscribe(topic)
	defer sub.Cancel()

	msg := "hello world"
	b.Publish(topic, msg)

	select {
	case received := <-sub.Chan:
		if received != msg {
			t.Fatalf("expected %q, got %q", msg, received)
		}
	case <-time.After(1 * time.Second):
		t.Fatal("did not receive message in time")
	}
}

func TestBrokerCancel(t *testing.T) {
	b := New()
	topic := "cancel_topic"
	sub := b.Subscribe(topic)
	sub.Cancel()

	b.Publish(topic, "should not be received")

	select {
	case _, ok := <-sub.Chan:
		if ok {
			t.Fatal("expected closed channel")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("timeout waiting for channel to close")
	}
}

