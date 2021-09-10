package observe

func NewSimpleObserver(id, event string, notify func(value Value)) SimpleObserver {
	return SimpleObserver{
		id:     id,
		event:  event,
		Handle: notify,
	}
}

type SimpleObserver struct {
	id     string
	event  string
	Handle func(val Value)
}

type AsyncSimpleObserver SimpleObserver

func (o SimpleObserver) ID() string {
	return o.id
}

func (o SimpleObserver) Event() string {
	return o.event
}

func (o SimpleObserver) Notify(val Value) {
	o.Handle(val)
}

func (o AsyncSimpleObserver) ID() string {
	return o.id
}

func (o AsyncSimpleObserver) Event() string {
	return o.event
}

func (o AsyncSimpleObserver) Notify(val Value) {
	go o.Handle(val)
}
