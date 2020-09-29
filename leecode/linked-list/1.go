package linked_list

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	tempNode := node.Next
	node.Val = tempNode.Val
	node.Next = tempNode.Next
	tempNode.Next = nil
}
