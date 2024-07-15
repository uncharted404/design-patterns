package singleton

type Singleton struct {
	name string
}

func NewSingleton(name string) *Singleton {
	return &Singleton{name: name}
}

func (s *Singleton) SayHello() {
	println(s)
	println("hello ", s.name)
}
