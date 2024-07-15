package simple_factory

import (
	impl2 "design-patterns/creational_patterns/simple_factory/impl"
)

const (
	AnimalDog = iota
	AnimalCat
)

type IAnimal interface {
	Eat()
}

func GetAnimal(animalName int) IAnimal {
	switch animalName {
	case AnimalDog:
		return &impl2.Dog{}
	case AnimalCat:
		return &impl2.Cat{}
	}
	return nil
}
