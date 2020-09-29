package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	current, pt := head, head
	cnt := 0
	for current.Next != nil{
		cnt++
		if cnt > n {
			pt = pt.Next
		}
		current = current.Next
	}

	if pt.Next == nil {
		return nil
	}
	if cnt == n-1{
		return pt.Next
	}

	pt.Next = pt.Next.Next
	return head
}