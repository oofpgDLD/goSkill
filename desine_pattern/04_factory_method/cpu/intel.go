package cpu

import "fmt"

type IntelFactory struct {
}

func (t *IntelFactory) Create() Cpu{
	return &IntelCpu{
	}
}

type IntelCpu struct {
}

func (t *IntelCpu)Open() {
	fmt.Println("amd cpu started")
}

func (t *IntelCpu)Close() {
	fmt.Println("amd cpu closed")
}