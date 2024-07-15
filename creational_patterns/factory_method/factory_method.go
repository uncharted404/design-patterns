package factory_method

type IAnimal interface {
	Eat()
}

type IAnimalFactory interface {
	CreateAnimal() IAnimal
}
