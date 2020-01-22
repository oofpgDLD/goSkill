package _4_factory_method

import (
	"errors"
	"fmt"
	"skillTest/desine_pattern/04_factory_method/cpu"
	"skillTest/desine_pattern/04_factory_method/desk"
	"skillTest/desine_pattern/04_factory_method/memory"
)

type Computer struct {
	powered bool
	mem memory.Mem
	cpu cpu.Cpu
	desk desk.Desk
}

func NewComputer(mem memory.Mem, cpu cpu.Cpu ,desk desk.Desk) *Computer{
	return &Computer{
		mem: mem,
		cpu: cpu,
		desk: desk,
	}
}

func (t *Computer) Run()error{
	if !t.powered {
		return errors.New("未通电，无法启动")
	}

	t.cpu.Open()
	t.mem.Open()
	t.desk.Open()
	return nil
}

func (t *Computer) Stop() error{
	if !t.powered {
		return errors.New("未通电，早已关闭")
	}

	t.desk.Close()
	t.mem.Close()
	t.cpu.Close()
	return nil
}

func (t *Computer) GetPower(voltage int) {
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