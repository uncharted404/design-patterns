package observer

type IObserver interface {
	Handle()
}

type ISubject interface {
	Register(IObserver)
	Notify()
}
