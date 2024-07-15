package observer

type Subject struct {
	observers []IObserver
}

func NewSubject() *Subject {
	return &Subject{}
}

func (s *Subject) Register(observer IObserver) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Notify() {
	for _, o := range s.observers {
		o.Handle()
	}
}

type Observer struct {
	name string
}

func NewObserver(name string) *Observer {
	return &Observer{name: name}
}

func (o *Observer) Handle() {
	println(o.name, "接到通知，正在处理")
}
