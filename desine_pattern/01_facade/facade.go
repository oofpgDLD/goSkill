package _1_facade

import (
	"errors"
	"fmt"
	"skillTest/desine_pattern/01_facade/cpu"
	"skillTest/desine_pattern/01_facade/desk"
	"skillTest/desine_pattern/01_facade/memory"
)

//外观模式

func NewMyComputer() *MyComputer{
	return &MyComputer{
		cpu: cpu.NewCpu("amd"),
		mem : memory.NewMemory("Kingston"),
		desk: desk.NewDesk("seagate"),
	}
}

type MyComputer struct {
	powered bool
	cpu cpu.Cpu
	mem memory.Memory
	desk desk.Desk
}

func (t *MyComputer) Run() error{
	if !t.powered {
		return errors.New("未通电，无法启动")
	}
	t.cpu.Start()
	t.mem.Start()
	t.desk.Start()
	return nil
}

func (t *MyComputer) Stop() error{
	if !t.powered {
		return errors.New("未通电，早已关闭")
	}
	t.cpu.Close()
	t.mem.Close()
	t.desk.Close()
	return nil
}

func (t *MyComputer) GetPower(voltage int) {
	if voltage > 5 {
		fmt.Println("电压过高，你的电脑炸了！")
		return
	}

	if voltage < 5 {
		fmt.Println("电压不足！")
		return
	}

	if voltage == 5 {
		t.powered = true
		return
	}
}