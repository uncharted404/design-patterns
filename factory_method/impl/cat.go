package impl

import "design-patterns/factory_method"

type Cat struct{}

func (d *Cat) Eat() {
	println("猫吃小鱼")
}

type CatFactory struct {
}

func NewCatFactory() factory_method.IAnimalFactory {
	return &CatFactory{}
}

func (f *CatFactory) CreateAnimal() factory_method.IAnimal {
	return &Cat{}
}
