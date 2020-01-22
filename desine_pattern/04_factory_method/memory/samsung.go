package memory

import "fmt"

type SamsungFactory struct {
}

func (t *SamsungFactory) Create() Mem{
	return &Samsung{
	}
}

type Samsung struct {
}

func (t *Samsung)Open() {
	fmt.Println("Samsung memory started")
}

func (t *Samsung)Close() {
	fmt.Println("Samsung memory closed")
}