package _02

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return add(l1,l2, 0)
}


func add(l1 *ListNode, l2 *ListNode, sub int) *ListNode {
	if l1 == nil && l2 == nil {
		if sub != 0 {
			return &ListNode{
				Val: sub,
				Next: nil,
			}
		}
		return nil
	}
	//
	vl1 := 0
	var pl1 *ListNode
	if l1 != nil {
		vl1 = l1.Val
		pl1 = l1.Next
	}

	vl2 := 0
	var pl2 *ListNode
	if l2 != nil {
		vl2 = l2.Val
		pl2 = l2.Next
	}

	v := vl1 + vl2 + sub
	carry := v / 10
	next := add(pl1, pl2, carry)
	return &ListNode{
		Val: v%10,
		Next: next,
	}
}