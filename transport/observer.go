package transport

import (
	"sync"
	"sync/atomic"
)

type Observer interface {
	Name() string           // on what event it should notify
	Update(val Event) error // notification callback
}

type Publisher struct {
	mx        sync.Mutex
	guid      *uint64
	observers map[uint64]Observer
}

func NewPublisher() *Publisher {
	var uid uint64 = 0
	return &Publisher{
		guid:      &uid,
		mx:        sync.Mutex{},
		observers: map[uint64]Observer{},
	}
}

// if event is empty then event broadcasting to all observers
func (o *Publisher) Broadcast(val Event) error {
	o.mx.Lock()
	defer o.mx.Unlock()
	for _, e := range o.observers {
		if err := e.Update(val); err != nil {
			return err
		}
	}
	return nil
}

// if Observer.Event == '*' then this Observer handles any events
func (o *Publisher) Notify(name string, val Event) error {
	o.mx.Lock()
	defer o.mx.Unlock()
	for _, e := range o.observers {
		switch e.Name() {
		case "*", name:
			if err := e.Update(val); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *Publisher) Register(val Observer) func() {
	o.mx.Lock()
	defer o.mx.Unlock()
	var uid = atomic.AddUint64(o.guid, 1)
	o.observers[uid] = val
	return func() {
		o.mx.Lock()
		defer o.mx.Unlock()
		delete(o.observers, uid)
	}
}

func NewSimpleObserver(name string, update func(value Event) error) SimpleObserver {
	return SimpleObserver{
		name:   name,
		update: update,
	}
}

type SimpleObserver struct {
	name   string
	update func(val Event) error
}

func (o SimpleObserver) Name() string {
	return o.name
}

func (o SimpleObserver) Update(val Event) error {
	return o.update(val)
}
