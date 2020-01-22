package cpu

import "fmt"

type AmdFactory struct {
}

func (t *AmdFactory) Create() Cpu{
	return &AmdCpu{
	}
}

type AmdCpu struct {
}

func (t *AmdCpu)Open() {
	fmt.Println("amd cpu started")
}

func (t *AmdCpu)Close() {
	fmt.Println("amd cpu closed")
}
