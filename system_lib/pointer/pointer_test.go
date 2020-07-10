package pointer

import (
	"testing"
	"unsafe"
)

type s struct {
	name string
}

func funcParam(in string, out *s) error{
	out = &s{name:in}
	return nil
}

func Test_UnsafePointer(t *testing.T) {
	a := 1

	t.Log(uintptr(a))
	t.Log(uintptr(unsafe.Pointer(&a)))
}

//go的函数的参数没有引用类型
//out不能在里面直接重新分配新的地址空间，只能改变原有out指向空间的值
func TestPointer_funcParam(t *testing.T) {
	var ret *s
	if err := funcParam("test", ret); err != nil {
		t.Error(err)
		return
	}
	t.Log(ret)
}

//逃逸
func TestRunAway(t *testing.T) {
	/*var f = func(x int) func() int{
		x++
		return func() int {
			return  x
		}
	}*/
	const N = 10
	var f = func() func() int{
		x := 0
		return func() int {
			x++
			return x
		}
	}
	ret := 0
	ff := f()
	for i:= 0; i < N ;i++ {
		ret = ff()
		t.Log(ret)
	}
}