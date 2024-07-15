package impl

import (
	"design-patterns/creational_patterns/abstract_factory"
	"fmt"
)

type ProductA1 struct {
	info string
}

func (p *ProductA1) ShowInfo() {
	println(p.info)
}

type ProductB1 struct {
	info string
}

func (p *ProductB1) ShowInfo() {
	println(p.info)
}

type Factory1 struct {
	name string
}

func NewFactory1() *Factory1 {
	return &Factory1{name: "Factory1"}
}

func (f *Factory1) CreateProductA() abstract_factory.IProductA {
	return &ProductA1{info: fmt.Sprintf("从%s生产的ProductA", f.name)}
}

func (f *Factory1) CreateProductB() abstract_factory.IProductB {
	return &ProductB1{info: fmt.Sprintf("从%s生产的ProductB", f.name)}
}
