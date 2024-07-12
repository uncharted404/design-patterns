package abstract_factory

type IFactory interface {
	CreateProductA() IProductA
	CreateProductB() IProductB
}

type IProductA interface {
	ShowInfo()
}

type IProductB interface {
	ShowInfo()
}
