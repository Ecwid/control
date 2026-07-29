package transport

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func newTestTransport(t *testing.T) *Transport {
	t.Helper()

	upgrader := websocket.Upgrader{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Errorf("upgrade failed: %v", err)
			return
		}
		defer conn.Close()

		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				return
			}
		}
	}))
	t.Cleanup(server.Close)

	wsURL := "ws" + server.URL[len("http"):]
	transport, err := Dial(context.Background(), websocket.Dialer{}, wsURL, nil)
	if err != nil {
		t.Fatalf("dial failed: %v", err)
	}
	t.Cleanup(func() {
		_ = transport.Close()
	})

	return transport
}

func TestSubscriberHubPublishFiltersBySession(t *testing.T) {
	hub := newSubscriberHub()

	all, cancelAll := hub.subscribe("", 1)
	defer cancelAll()

	s1, cancelS1 := hub.subscribe("s1", 1)
	defer cancelS1()

	hub.publish(Message{SessionID: "s1", Method: "evt"})

	select {
	case <-all:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("global subscriber did not receive message")
	}

	select {
	case <-s1:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("session subscriber did not receive message")
	}
}

func TestSubscriberHubSlowSubscriberDoesNotBlock(t *testing.T) {
	hub := newSubscriberHub()

	slow, cancelSlow := hub.subscribe("", 1)
	defer cancelSlow()
	_, _ = slow, cancelSlow

	start := time.Now()
	for i := 0; i < 5000; i++ {
		hub.publish(Message{Method: "evt"})
	}
	if time.Since(start) > 200*time.Millisecond {
		t.Fatalf("publish appears blocking for slow subscriber: %v", time.Since(start))
	}
}

func TestSubscriberHubUnsubscribeIsSafe(t *testing.T) {
	hub := newSubscriberHub()

	ch, cancel := hub.subscribe("", 1)
	cancel()
	cancel()

	select {
	case _, ok := <-ch:
		if ok {
			t.Fatal("channel must be closed after unsubscribe")
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("unsubscribe did not close channel")
	}
}

func TestTransportCloseClosesSubscriberChannels(t *testing.T) {
	transport := newTestTransport(t)

	ch, cancel := transport.Subscribe("", 1)
	defer cancel()

	if err := transport.Close(); err != nil {
		t.Fatalf("close failed: %v", err)
	}

	select {
	case _, ok := <-ch:
		if ok {
			t.Fatal("subscriber channel must be closed on transport close")
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("subscriber channel was not closed on transport close")
	}
}
