package desk

import "fmt"

func NewDesk(tp string) Desk{
	switch tp {
	case "seagate":
		return &Seagate{}
	}
	return nil
}

type Desk interface {
	Start()
	Close()
}

type Seagate struct {
}

func (t *Seagate)Start() {
	fmt.Println("希捷硬盘启动")
}

func (t *Seagate)Close() {
	fmt.Println("希捷硬盘关闭")
}