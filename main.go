package main 

import (
	"fmt"
)

var (
	a string
)

func init() {
	a = "this is main init"
}

type TestType struct {
	b string
}

func (t *TestType) init() {
	t.b = "this struct init"
}

func test () (a ,b int) {
	a = 1
	b = 2
	return
}

func main(){
	id := "52BD06"

	if len(id) >= 8 {
		fmt.Println(id[0:8])
	} else {
		fmt.Println(id[0:])
	}

	a,b:= test()
	fmt.Println(a,b)
}