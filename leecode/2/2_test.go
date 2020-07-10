package _02

import (
	"testing"
)

func Test_Way(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
				Next: nil,
			},
		},
	}

	ret := addTwoNumbers(l1, l2)
	p := ret
	for p != nil{
		t.Log(p.Val)
		p = p.Next
	}
}