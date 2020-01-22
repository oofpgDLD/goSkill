package memory

import "fmt"

type KingstonFactory struct {
}

func (t *KingstonFactory) Create() Mem{
	return &Kingston{
	}
}

type Kingston struct {
}

func (t *Kingston)Open() {
	fmt.Println("Kingston memory started")
}

func (t *Kingston)Close() {
	fmt.Println("Kingston memory closed")
}

