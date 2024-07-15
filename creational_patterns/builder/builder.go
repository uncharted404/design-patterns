package builder

import (
	"fmt"
	"strings"
)

type PcBuilder struct {
	cpu       string
	gpu       string
	mainBoard string
}

func NewPcBuilder() *PcBuilder {
	return &PcBuilder{}
}

func (p *PcBuilder) SetCPU(cpu string) *PcBuilder {
	p.cpu = cpu
	return p
}

func (p *PcBuilder) SetGPU(gpu string) *PcBuilder {
	p.gpu = gpu
	return p
}

func (p *PcBuilder) SetMainBoard(mainBoard string) *PcBuilder {
	p.mainBoard = mainBoard
	return p
}

func (p *PcBuilder) Build() *Pc {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("CPU选择：%s\n", p.cpu))
	s.WriteString(fmt.Sprintf("GPU选择：%s\n", p.gpu))
	s.WriteString(fmt.Sprintf("主板选择：%s\n", p.mainBoard))
	s.WriteString("组装完成！！\n")
	return newPc(s.String())
}
