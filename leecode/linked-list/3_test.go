package linked_list

import (
	"os"
	"testing"
)

var head *ListNode

func TestMain(m *testing.M) {
	collection := []int{5,4,3,2,1}
	for _, item := range collection {
		node := &ListNode{
			Val: item,
			Next: head,
		}
		head = node
	}
	os.Exit(m.Run())
}

func Test_reverseList2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{head: head}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := reverseList2(tt.args.head)
			for pt != nil{
				t.Log(pt.Val)
				pt = pt.Next
			}
		})
	}
}

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{head: head}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := reverseList(tt.args.head)
			for pt != nil{
				t.Log(pt.Val)
				pt = pt.Next
			}
		})
	}
}