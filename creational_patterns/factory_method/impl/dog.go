package impl

import (
	"design-patterns/creational_patterns/factory_method"
)

type Dog struct{}

func (d *Dog) Eat() {
	println("狗吃骨头")
}

type DogFactory struct {
}

func NewDogFactory() factory_method.IAnimalFactory {
	return &DogFactory{}
}

func (f *DogFactory) CreateAnimal() factory_method.IAnimal {
	return &Dog{}
}
