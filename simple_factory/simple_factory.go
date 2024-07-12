package simple_factory

import "design-patterns/simple_factory/impl"

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
		return &impl.Dog{}
	case AnimalCat:
		return &impl.Cat{}
	}
	return nil
}
