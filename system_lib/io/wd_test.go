package io

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	wd := &WallDoor{
		buffer: make([]byte, 200),
		ch: make(chan byte, 1),
	}
	g := &NumberStacker{}
	for i:=0;i < 120;i++ {
		wd.generate(g)
	}
	t.Log(wd.buffer)
	t.Log(wd.w)
}

func TestFilling(t *testing.T) {
	wd := &WallDoor{
		buffer: make([]byte, 200),
		ch: make(chan byte, 1),
	}
	g := &NumberStacker{}
	for i:=0;i < 120;i++ {
		wd.generate(g)
	}
	for i:=0;i < 200;i++ {
		wd.filling()
	}
	t.Log(wd.buffer)
	t.Log(wd.w)
}