package cpu

import "fmt"

func NewCpu(tp string) Cpu{
	switch tp {
	case "amd":
		return &AmdCpu{}
	}
	return nil
}

type Cpu interface {
	Start()
	Close()
}

type AmdCpu struct {
}

func (t *AmdCpu)Start() {
	fmt.Println("amd cpu started")
}

func (t *AmdCpu)Close() {
	fmt.Println("amd cpu closed")
}