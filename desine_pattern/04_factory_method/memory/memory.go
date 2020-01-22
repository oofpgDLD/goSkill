package memory

type IFactoryMem interface {
	Create() Mem
}

type Mem interface {
	Open()
	Close()
}