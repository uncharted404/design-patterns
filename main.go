package main

import (
	"design-patterns/behavioral_patterns/chain_of_responsibility"
	"design-patterns/behavioral_patterns/observer"
	afimpl "design-patterns/creational_patterns/abstract_factory/impl"
	"design-patterns/creational_patterns/builder"
	amimpl "design-patterns/creational_patterns/factory_method/impl"
	"design-patterns/creational_patterns/prototype"
	"design-patterns/creational_patterns/simple_factory"
	"design-patterns/creational_patterns/singleton/hungry"
	"design-patterns/creational_patterns/singleton/lazy"
	"design-patterns/structural_patterns/adapter"
	"design-patterns/structural_patterns/bridge"
	"design-patterns/structural_patterns/proxy"
	"fmt"
)

func AbstractFactory() {
	factory1 := afimpl.NewFactory1()
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
	catFactory := amimpl.NewCatFactory()
	cat := catFactory.CreateAnimal()
	cat.Eat()

	dogFactory := amimpl.NewDogFactory()
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
	byCas := hungry.GetInstanceByCas()

	byInit.SayHello()
	byVar.SayHello()

	byMutex.SayHello()
	byOnce.SayHello()
	byCas.SayHello()
}

func Adapter() {
	httpClient := adapter.NewHttpClient()
	client := adapter.NewHttpClientAdapter(httpClient)
	client.Request()
}

func Bridge() {
	smsNotification := bridge.NewSmsNotification()
	emailNotification := bridge.NewEmailNotification()

	smsNotification.Notice("sms info")
	emailNotification.Notice("email info")
}

func ChainOfResponsibility() {
	router := chain_of_responsibility.NewRouter()

	router.Use(func(ctx *chain_of_responsibility.HttpCtx) {
		println("1111111")
		ctx.Next()
		println("2222222")
	})

	router.Use(func(ctx *chain_of_responsibility.HttpCtx) {
		println("3333333")
		ctx.Next()
		println("4444444")
	})

	router.Get("/get/hello", func(ctx *chain_of_responsibility.HttpCtx) {
		ctx.Resp(200, map[string]interface{}{
			"resp": "hello",
		})
	})

	router.Run()
}

func Observer() {
	o1 := observer.NewObserver("1")
	o2 := observer.NewObserver("2")
	o3 := observer.NewObserver("3")

	subject := observer.NewSubject()
	subject.Register(o1)
	subject.Register(o2)
	subject.Register(o3)

	subject.Notify()
}

func Proxy() {
	clientProxy := proxy.NewRpcClientProxy()
	resp := clientProxy.Request()
	println(resp)
}

func main() {
	//AbstractFactory()
	//Builder()
	//FactoryMethod()
	//Prototype()
	//SimpleFactory()
	//Singleton()

	//Adapter()
	//Bridge()
	//Proxy()

	//ChainOfResponsibility()
	Observer()
}

func Test() {

}
