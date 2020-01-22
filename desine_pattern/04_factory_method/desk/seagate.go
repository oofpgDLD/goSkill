package desk

import "fmt"

type SeagateFactory struct {

}

func (t *SeagateFactory) Create() Desk{
	return &Seagate{

	}
}

type Seagate struct {

}

func (t *Seagate)Open() {
	fmt.Println("希捷硬盘启动")
}

func (t *Seagate)Close() {
	fmt.Println("希捷硬盘关闭")
}