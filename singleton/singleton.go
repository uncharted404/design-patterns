package singleton

type Singleton struct {
}

func NewSingleton() *Singleton {
	return &Singleton{}
}

func (s *Singleton) SayHello() {
	println("Hello")
}
