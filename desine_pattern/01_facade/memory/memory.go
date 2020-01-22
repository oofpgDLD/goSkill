package memory

import "fmt"

func NewMemory(tp string) Memory{
	switch tp {
	case "Kingston":
		return &Kingston{}
	}
	return nil
}

type Memory interface {
	Start()
	Close()
}

type Kingston struct {
}

func (t *Kingston) Start() {
	fmt.Println("Kingston memory started")
}

func (t *Kingston) Close() {
	fmt.Println("Kingston memory closed")
}