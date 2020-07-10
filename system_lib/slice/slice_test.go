package slice

import (
	"testing"
	"unsafe"
)

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func Test_SliceCopy(t *testing.T) {
	s1 := []int{1,2,3}
	s2 := s1


	t.Log("s1 address", unsafe.Pointer(&s1))
	t.Log("s2 address", unsafe.Pointer(&s2))

	s1[0] = 4
	t.Logf("s1 第0个元素改为4后, s2第0个元素为%v",s2[0])
	s3 := s2[:2]
	t.Logf("s3:%v,len:%v,cap:%v", s3, len(s3), cap(s3))
	t.Logf("s1:%v,len:%v,cap:%v", s1, len(s1), cap(s1))
	t.Logf("s2:%v,len:%v,cap:%v", s2, len(s2), cap(s2))

	s1[0] = 5
	t.Logf("s1第0个元素改为5后，s2[0]=%v.s3[0]=%v", s2[0], s3[0])
}


func Test_DeepCopy(t *testing.T) {
	s1 := []int{1,2,3}
	s2 := make([]int,3)
	s2 = s1

	t.Log("s1 address", unsafe.Pointer(&s1))
	t.Log("s2 address", unsafe.Pointer(&s2))

	s1[0] = 4
	t.Logf("s1 第0个元素改为4后, s2第0个元素为%v",s2[0])

	//s3为深拷贝
	s3 := make([]int,3)
	copy(s3, s2)

	t.Logf("s3:%v,len:%v,cap:%v", s3, len(s3), cap(s3))
	t.Logf("s1:%v,len:%v,cap:%v", s1, len(s1), cap(s1))
	t.Logf("s2:%v,len:%v,cap:%v", s2, len(s2), cap(s2))

	s1[0] = 5
	t.Logf("s1第0个元素改为5后，s2[0]=%v.s3[0]=%v", s2[0], s3[0])
}

func Test_SliceFunc(t *testing.T) {
	s1 := []int{1,2,3}
	s2 := make([]int,3)
	s2 = s1

	t.Log("s1 address", unsafe.Pointer(&s1))
	t.Log("s2 address", unsafe.Pointer(&s2))

	s1[0] = 4
	t.Logf("s1第0个元素改为4后, s2第0个元素为%v",s2[0])

	//函数参数传递
	func (s []int) {
		s[0] = 5
	}(s1)
	t.Logf("函数中s第0个元素改为5后，s1[0]=%v,s2[0]=%v",s1[0], s2[0])
}

func Test_SliceAppend(t *testing.T) {
	s1 := []int{1,2,3}
	s2 := make([]int,3)
	s2 = s1

	t.Log("s1 address", unsafe.Pointer(&s1))
	t.Log("s2 address", unsafe.Pointer(&s2))

	s1[0] = 4
	t.Logf("s1第0个元素改为4后, s2第0个元素为%v",s2[0])

	//函数参数传递
	s3 := func (s []int) []int{
		s[0] = 5
		s = append(s, 4)
		s[1] = 8
		return s
	}(s1)
	t.Logf("s3:%v,len:%v,cap:%v", s3, len(s3), cap(s3))
	t.Logf("函数中s第0个元素改为5后，s1[0]=%v,s2[0]=%v,s3[0]=%v",s1[0], s2[0], s3[0])
	t.Logf("函数中s append后第1个元素改为8后，s1[1]=%v,s2[1]=%v,s3[1]=%v",s1[1], s2[1], s3[1])
}

func Test_SliceAppend2(t *testing.T) {
	s1 := []int{1,2,3}
	s2 := make([]int,3)
	s2 = s1

	t.Log("s1 address", unsafe.Pointer(&s1))
	t.Log("s2 address", unsafe.Pointer(&s2))

	s1[0] = 4
	t.Logf("s1第0个元素改为4后, s2第0个元素为%v",s2[0])

	//函数参数传递
	s3 := func (s []int) []int{
		s[0] = 5
		s = append(s, 4)
		s[1] = 8
		return s
	}(s1)
	t.Logf("函数中s第0个元素改为5后，s1[0]=%v,s2[0]=%v,s3[0]=%v",s1[0], s2[0], s3[0])
	t.Logf("函数中s append后第1个元素改为8后，s1[1]=%v,s2[1]=%v,s3[1]=%v",s1[1], s2[1], s3[1])

	func (s []int) []int{
		s[0] = -1
		s = append(s, 4)
		s[1] = -8
		return s
	}(s3)

	t.Logf("函数中s第0个元素改为-1后，s1[0]=%v,s2[0]=%v,s3[0]=%v",s1[0], s2[0], s3[0])
	t.Logf("函数中s append后第1个元素改为-8后，s1[1]=%v,s2[1]=%v,s3[1]=%v",s1[1], s2[1], s3[1])
}