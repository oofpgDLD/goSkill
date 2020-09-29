package linked_list

var pt *ListNode

func checkNode(head *ListNode) bool {
	if head == nil {
		return true
	}
	b := checkNode(head.Next)
	if !b {
		return false
	}
	if head.Next != nil {
		pt = pt.Next
	}
	return pt.Val == head.Val
}

func isPalindrome(head *ListNode) bool {
	pt = head
	return checkNode(head)
}