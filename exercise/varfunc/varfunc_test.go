package varfunc

import (
	"fmt"
	"testing"
)

var h1 func(string)



func Test_FuncHandle(t *testing.T) {
	h1 := func(s string) {
		fmt.Println(s)
	}

	t.Log(h1())
}