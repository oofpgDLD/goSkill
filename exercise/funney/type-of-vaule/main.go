package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	name string
}

func main() {

	var o interface{}
	ms := MyStruct{}
	o = ms
	fmt.Println(reflect.TypeOf(1))
	fmt.Println(reflect.TypeOf(9.5))
	fmt.Println(reflect.TypeOf("Just a String"))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf(ms))

	switch t := o.(type) {
	case MyStruct:
		fmt.Println(t)
	default:
		fmt.Println("未识别的类型")
	}
}