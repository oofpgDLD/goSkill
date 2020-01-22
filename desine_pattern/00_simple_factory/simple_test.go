package _0_simple_factory

import "testing"

func Test_01(t *testing.T) {
	c := NewCar("bmw")
	if c == nil {
		t.Error("car not find")
		return
	}
	c.Drive()
	return
}

func Test_02(t *testing.T) {
	list := []string{"bmw","benz","audi"}
	for _,name := range list{
		c := NewCar(name)
		if c == nil {
			t.Errorf("%v car not find", name)
		}else {
			c.Drive()
		}
	}
	return
}