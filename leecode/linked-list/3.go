package linked_list

//迭代法
func reverseList(head *ListNode) *ListNode {
	var pt,temp *ListNode
	for head != nil {
		temp = head.Next
		head.Next = pt
		pt = head
		head = temp
	}
	return pt
}

//递归法
func reverseList2(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	pt := reverseList2(head.Next)
	pt.Next = head
	return head
}