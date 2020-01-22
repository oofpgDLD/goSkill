package _6_builder

import (
	"fmt"
	"skillTest/desine_pattern/04_factory_method"
	"skillTest/desine_pattern/04_factory_method/cpu"
	"skillTest/desine_pattern/04_factory_method/desk"
	"skillTest/desine_pattern/04_factory_method/memory"
)

type BrandComputer interface {
	SetName(string)
	GetName() string
	Construct(*_4_factory_method.Computer)
}

type Lenovo struct {
	Name string
	base *_4_factory_method.Computer
}

func (t *Lenovo) SetName(name string) {
	t.Name = name
}

func (t *Lenovo) GetName() string {
	return t.Name
}

func (t *Lenovo) Construct(b *_4_factory_method.Computer) {
	t.base = b
}

func (t *Lenovo) Run() error{
	fmt.Println(t.GetName())
	return t.base.Run()
}

type Acer struct {
	Name string
	base *_4_factory_method.Computer
}

func (t *Acer) SetName(name string) {
	t.Name = name
}

func (t *Acer) GetName() string {
	return t.Name
}

func (t *Acer) Construct(b *_4_factory_method.Computer) {
	t.base = b
}

func (t *Acer) Run() error{
	fmt.Println(t.GetName())
	return t.base.Run()
}

type Director struct {
	builder BrandComputer
}

// NewDirector ...
func NewDirector(builder BrandComputer) *Director {
	return &Director{
		builder: builder,
	}
}

//Construct Product
func (d *Director) Construct() {
	var cpufact cpu.IFactoryCpu
	var memfact memory.IFactoryMem
	var deskfact desk.IFactoryDesk

	builder := d.builder
	switch builder.(type) {
	case *Lenovo:
		builder.SetName("Lenovo")
		cpufact = &cpu.IntelFactory{}
		memfact = &memory.KingstonFactory{}
		deskfact = &desk.WdFactory{}
	case *Acer:
		builder.SetName("Acer")
		cpufact = &cpu.AmdFactory{}
		memfact = &memory.SamsungFactory{}
		deskfact = &desk.SeagateFactory{}
	}
	builder.Construct(_4_factory_method.NewComputer(memfact.Create(), cpufact.Create(), deskfact.Create()))
}