package _1_facade

import (
	"fmt"
	"testing"
	"time"
)

func Test_MyComputer(t *testing.T) {
	c := NewMyComputer()
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