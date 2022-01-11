package transport

import (
	"sync"
)

type Observer interface {
	ID() string       // unique observer's id, attaching and detaching by this id
	Event() string    // on what event it should notified
	Update(val Event) // notification callback
}

type Publisher struct {
	mx        sync.Mutex
	observers map[string]Observer
}

func NewPublisher() *Publisher {
	return &Publisher{
		mx:        sync.Mutex{},
		observers: map[string]Observer{},
	}
}

// if event is empty then event broadcasting to all observers
// if Observer.Event == '*' then this Observer handles any events
func (o *Publisher) Notify(event string, val Event) {
	o.mx.Lock()
	defer o.mx.Unlock()
	for _, e := range o.observers {
		if e.Event() == "*" || event == "" || e.Event() == event {
			e.Update(val)
		}
	}
}

func (o *Publisher) Register(val Observer) {
	o.mx.Lock()
	defer o.mx.Unlock()
	o.observers[val.ID()] = val
}

func (o *Publisher) Unregister(val Observer) {
	o.mx.Lock()
	defer o.mx.Unlock()
	delete(o.observers, val.ID())
}

func NewSimpleObserver(id, event string, update func(value Event)) SimpleObserver {
	return SimpleObserver{
		id:     id,
		event:  event,
		update: update,
	}
}

type SimpleObserver struct {
	id     string
	event  string
	update func(val Event)
}

func (o SimpleObserver) ID() string {
	return o.id
}

func (o SimpleObserver) Event() string {
	return o.event
}

func (o SimpleObserver) Update(val Event) {
	o.update(val)
}
