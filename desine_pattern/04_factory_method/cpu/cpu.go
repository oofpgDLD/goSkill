package cpu

type IFactoryCpu interface {
	Create() Cpu
}

type Cpu interface {
	Open()
	Close()
}