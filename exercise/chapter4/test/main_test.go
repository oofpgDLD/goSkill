package test

import "testing"

func Test_Array(t *testing.T) {
	r := [...]int{99:-1}
	t.Log(r)
}

func Test_ArrayEqual(t *testing.T) {
	r1 := [...]int{1,2,3}
	r2 := [...]int{1,2,3}
	t.Log(r1 == r2)
}

func Test_SliceAppend(t *testing.T) {
	a1 := [...]int{1,2,3,4,5,6}
	s1 := a1[:1]
	s2 := append(s1, 100)
	s3 := s1[:2]
	t.Log(s1, len(s1), cap(s1))
	t.Log(s2, len(s2), cap(s2))
	t.Log(s3, len(s3), cap(s3))
}
