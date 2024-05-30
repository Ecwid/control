package cdp

import (
	"sync"
)

var BrokerChannelSize = 50000

type subscriber struct {
	sessionID string
	channel   chan Message
}

type broker struct {
	cancel   chan struct{}
	messages chan Message
	sub      chan subscriber
	unsub    chan chan Message
	lock     *sync.Mutex
}

func makeBroker() broker {
	return broker{
		cancel:   make(chan struct{}),
		messages: make(chan Message),
		sub:      make(chan subscriber),
		unsub:    make(chan chan Message),
		lock:     &sync.Mutex{},
	}
}

func (b broker) run() {
	var value = map[chan Message]subscriber{}
	for {
		select {

		case sub := <-b.sub:
			value[sub.channel] = sub

		case channel := <-b.unsub:
			if _, ok := value[channel]; ok {
				delete(value, channel)
				close(channel)
			}

		case <-b.cancel:
			for msgCh := range value {
				close(msgCh)
			}
			close(b.sub)
			close(b.unsub)
			close(b.messages)
			return

		case message := <-b.messages:
			for _, subscriber := range value {
				if message.SessionID == "" || subscriber.sessionID == "" || message.SessionID == subscriber.sessionID {
					subscriber.channel <- message
				}
			}
		}
	}
}

func (b broker) subscribe(sessionID string) chan Message {
	b.lock.Lock()
	defer b.lock.Unlock()

	select {
	case <-b.cancel:
		return nil
	default:
		sub := subscriber{
			sessionID: sessionID,
			channel:   make(chan Message, BrokerChannelSize),
		}
		b.sub <- sub
		return sub.channel
	}
}

func (b broker) unsubscribe(value chan Message) {
	b.lock.Lock()
	select {
	case <-b.cancel:
	default:
		b.unsub <- value
	}
	b.lock.Unlock()
}

func (b broker) publish(msg Message) {
	b.messages <- msg
}

func (b broker) Cancel() {
	b.lock.Lock()
	close(b.cancel)
	b.lock.Unlock()
}
