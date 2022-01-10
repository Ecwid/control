package transport

import (
	"sync"
)

type Observer interface {
	Hash() string     // unique observer's id, attaching and detaching by this id
	Event() string    // on what event it should notified
	Update(val Event) // notification callback
}

type Publisher struct {
	mx        sync.Mutex
	observers []Observer
}

func NewPublisher() *Publisher {
	return &Publisher{
		mx:        sync.Mutex{},
		observers: make([]Observer, 0),
	}
}

// if event is empty then event broadcasting to all observers
// if Observer.Event == '*' then this Observer handles any events
func (o *Publisher) Notify(event string, val Event) {
	o.mx.Lock()
	defer o.mx.Unlock()
	for _, e := range o.observers {
		if (e.Event() == "*" || event == "" || e.Event() == event) && e.Update != nil {
			e.Update(val)
		}
	}
}

func (o *Publisher) Register(val Observer) {
	o.mx.Lock()
	defer o.mx.Unlock()
	o.observers = append(o.observers, val)
}

func (o *Publisher) Unregister(val Observer) {
	o.mx.Lock()
	defer o.mx.Unlock()
	for n, e := range o.observers {
		if e.Hash() == val.Hash() {
			tail := len(o.observers) - 1
			o.observers[n] = o.observers[tail]
			o.observers = o.observers[:tail]
			return
		}
	}
}

func NewSimpleObserver(hash, event string, update func(value Event)) SimpleObserver {
	return SimpleObserver{
		hash:   hash,
		event:  event,
		update: update,
	}
}

type SimpleObserver struct {
	hash   string
	event  string
	update func(val Event)
}

func (o SimpleObserver) Hash() string {
	return o.hash
}

func (o SimpleObserver) Event() string {
	return o.event
}

func (o SimpleObserver) Update(val Event) {
	o.update(val)
}
