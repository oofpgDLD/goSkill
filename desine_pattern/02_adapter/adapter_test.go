package _2_adapter

import (
	"fmt"
	"skillTest/desine_pattern/01_facade"
	"testing"
	"time"
)

func Test_Adapter(t *testing.T) {
	//家在中国
	ac := AC220{}
	//使用中国的适配器
	var adapter DC5Adapter
	adapter = &ChinaPowerAdapter{}
	//输出直流电压
	dc := adapter.OutputDC5(&ac)
	c := _1_facade.NewMyComputer()
	c.GetPower(dc)
	//运行电脑
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