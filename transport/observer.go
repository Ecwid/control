package transport

import (
	"sync"
)

type Observer interface {
	ID() string             // unique observer's id, attaching and detaching by this id
	Event() string          // on what event it should notify
	Update(val Event) error // notification callback
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
func (o *Publisher) Notify(event string, val Event) error {
	o.mx.Lock()
	defer o.mx.Unlock()
	for _, e := range o.observers {
		switch e.Event() {
		case "", "*", event:
			if err := e.Update(val); err != nil {
				return err
			}
		}
	}
	return nil
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

func NewSimpleObserver(id, event string, update func(value Event) error) SimpleObserver {
	return SimpleObserver{
		id:     id,
		event:  event,
		update: update,
	}
}

type SimpleObserver struct {
	id     string
	event  string
	update func(val Event) error
}

func (o SimpleObserver) ID() string {
	return o.id
}

func (o SimpleObserver) Event() string {
	return o.event
}

func (o SimpleObserver) Update(val Event) error {
	return o.update(val)
}
