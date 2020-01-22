package _0_simple_factory

import "fmt"

//简单工厂模式

type Car interface {
	Drive()
}

type BenzCar struct {
}

type AudiCar struct {
}

type BmwCar struct {
}

func (t *BenzCar)Drive() {
	fmt.Printf("奔驰车启动：%v","最大速度220km/h\n")
}

func (t *AudiCar)Drive() {
	fmt.Printf("奥迪车启动：%v","最大速度230km/h\n")
}

func (t *BmwCar)Drive() {
	fmt.Printf("宝马车启动：%v","最大速度200km/h\n")
}


func NewCar(tp string) Car{
	switch tp {
	case "bmw":
		return &BmwCar{}
	case "audi":
		return &AudiCar{}
	case "benz":
		return &BenzCar{}
	default:
		return nil
	}
}




