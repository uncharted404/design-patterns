package builder

type Pc struct {
	info string
}

func newPc(info string) *Pc {
	return &Pc{info: info}
}
func (p *Pc) ShowInfo() {
	println(p.info)
}
