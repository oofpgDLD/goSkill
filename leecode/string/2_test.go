package string

import "testing"

func TestReverse(t *testing.T) {
	a := -123
	ret := reverse(a)
	t.Log(ret)
}
