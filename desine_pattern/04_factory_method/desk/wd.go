package desk

import "fmt"

type WdFactory struct {

}

func (t *WdFactory) Create() Desk{
	return &Seagate{

	}
}

type Wd struct {

}

func (t *Wd) Open() {
	fmt.Println("西数硬盘启动")
}

func (t *Wd) Close() {
	fmt.Println("西数硬盘关闭")
}