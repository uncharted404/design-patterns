package main

import (
	impl2 "design-patterns/abstract_factory/impl"
	"design-patterns/builder"
	"design-patterns/factory_method/impl"
	"design-patterns/prototype"
	"design-patterns/simple_factory"
	"design-patterns/singleton/hungry"
	"design-patterns/singleton/lazy"
	"fmt"
)

func AbstractFactory() {
	factory1 := impl2.NewFactory1()
	productA := factory1.CreateProductA()
	productB := factory1.CreateProductB()
	productA.ShowInfo()
	productB.ShowInfo()
}

func Builder() {
	highPc := builder.NewPcBuilder().
		SetCPU("intel i9-13490F").
		SetGPU("华硕/七彩虹RTX4090-24G").
		SetMainBoard("七彩虹Colorful B760M WIFI").
		Build()

	lowPc := builder.NewPcBuilder().
		SetCPU("intel i3-13490F").
		SetGPU("华硕/七彩虹RTX4060-8G").
		SetMainBoard("七彩虹Colorful B760M WIFI").
		Build()

	highPc.ShowInfo()
	lowPc.ShowInfo()
}

func FactoryMethod() {
	catFactory := impl.NewCatFactory()
	cat := catFactory.CreateAnimal()
	cat.Eat()

	dogFactory := impl.NewDogFactory()
	dog := dogFactory.CreateAnimal()
	dog.Eat()
}

func Prototype() {
	p1 := prototype.NewPrototype("大黄", []string{"1", "2", "3"})
	p2 := p1.Clone().(*prototype.Prototype)
	p1.List = append(p1.List, "4")
	fmt.Printf("%+v\n%+v\n", p1, p2)
}

func SimpleFactory() {
	dog := simple_factory.GetAnimal(simple_factory.AnimalDog)
	cat := simple_factory.GetAnimal(simple_factory.AnimalCat)
	dog.Eat()
	cat.Eat()
}

func Singleton() {
	byInit := lazy.GetInstanceByInit()
	byVar := lazy.GetInstanceByVar()

	byMutex := hungry.GetInstanceByMutex()
	byOnce := hungry.GetInstanceByOnce()

	byInit.SayHello()
	byVar.SayHello()
	byMutex.SayHello()
	byOnce.SayHello()
}

func main() {
	AbstractFactory()
	//Builder()
	//FactoryMethod()
	//Prototype()
	//SimpleFactory()
	//Singleton()
}
