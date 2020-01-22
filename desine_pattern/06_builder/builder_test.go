package _6_builder

import (
	"fmt"
	"skillTest/desine_pattern/02_adapter"
	"testing"
	"time"
)

func Test_BuildLenovo(t *testing.T) {
	//家在中国
	ac := _2_adapter.AC220{}
	//使用中国的适配器
	var adapter _2_adapter.DC5Adapter
	adapter = &_2_adapter.ChinaPowerAdapter{}
	//输出直流电压
	dc := adapter.OutputDC5(&ac)

	c := &Lenovo{}
	director := NewDirector(c)
	director.Construct()
	c.base.GetPower(dc)

	err := c.Run()
	if err != nil {
		t.Log(err.Error())
		return
	}
	fmt.Println("using...")
	time.Sleep(time.Second * 5)
	err = c.base.Stop()
	if err != nil {
		t.Log(err.Error())
	}
}