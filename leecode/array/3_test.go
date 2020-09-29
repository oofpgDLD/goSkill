package array

import "testing"

func TestRotate(t *testing.T) {
	a := []int{1}
	rotate(a,0)
	t.Log(a)
}
