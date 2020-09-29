package array

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	a := []int{1,1,2}
	num := removeDuplicates(a)
	t.Log(num)
	t.Log(a)
}
