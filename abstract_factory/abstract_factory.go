package abstract_factory

type AbstractFactory interface {
	CreateProductA() ProductA
	CreateProductB() ProductB
}

type ProductA interface {
	ShowInfo()
}

type ProductB interface {
	ShowInfo()
}
