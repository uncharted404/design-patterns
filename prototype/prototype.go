package prototype

type IPrototype interface {
	Clone() IPrototype
}

type Prototype struct {
	Name string
	List []string
}

func NewPrototype(name string, list []string) *Prototype {
	return &Prototype{Name: name, List: list}
}

func (p *Prototype) Clone() IPrototype {
	newList := make([]string, len(p.List))
	copy(newList, p.List)
	return &Prototype{Name: p.Name, List: newList}
}
