package _4_factory_method

import (
	"fmt"
	"skillTest/desine_pattern/02_adapter"
	"skillTest/desine_pattern/04_factory_method/cpu"
	"skillTest/desine_pattern/04_factory_method/desk"
	"skillTest/desine_pattern/04_factory_method/memory"
	"testing"
	"time"
)

func Test_Factory_method(t *testing.T) {
	var cpufact cpu.IFactoryCpu
	var memfact memory.IFactoryMem
	var deskfact desk.IFactoryDesk

	cpufact = &cpu.IntelFactory{}
	memfact = &memory.KingstonFactory{}
	deskfact = &desk.WdFactory{}


	//家在中国
	ac := _2_adapter.AC110{}
	//使用中国的适配器
	var adapter _2_adapter.DC5Adapter
	adapter = &_2_adapter.ChinaPowerAdapter{}
	//输出直流电压
	dc := adapter.OutputDC5(&ac)

	c := NewComputer(memfact.Create(), cpufact.Create(), deskfact.Create())
	c.GetPower(dc)

	err := c.Run()
	if err != nil {
		t.Log(err.Error())
		return
	}
	fmt.Println("using...")
	time.Sleep(time.Second * 5)
	err = c.Stop()
	if err != nil {
		t.Log(err.Error())
	}
}