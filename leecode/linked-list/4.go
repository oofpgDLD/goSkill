package linked_list

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	pt1, pt2 := l1 ,l2
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		pt1, pt2 = l2 ,l1
	}
	hd := pt1
	for {
		if pt2 == nil {
			break
		}
		if pt1.Next == nil {
			pt1.Next = pt2
			break
		}
		if pt1.Next.Val < pt2.Val{
			//l1后移
			pt1 = pt1.Next
		} else {
			//l2后移
			node := pt2
			pt2 = pt2.Next
			//插入l2元素
			temp := pt1.Next
			pt1.Next = node
			node.Next = temp
		}
	}
	return hd
}