package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	/*s := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right:&TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 0,
				},
			},
		},
		Right:&TreeNode{
			Val: 5,
		},
	}

	t := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
		Right:&TreeNode{
			Val: 2,
		},
	}*/

	s := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right:&TreeNode{
			Val: 3,
		},
	}

	t := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}

	fmt.Println(isSubtree(s, t))
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil {
		return s == t
	}
	if s.Val != t.Val || !(isEqualTree(s.Left, t.Left) && isEqualTree(s.Right, t.Right)){
		//比较left
		if !isSubtree(s.Left, t){
			//比较right
			return isSubtree(s.Right, t)
		}
	}
	return true
}

func isEqualTree(s *TreeNode, t *TreeNode) bool{
	if s == nil || t == nil {
		return s == t
	}
	if s.Val != t.Val{
		return false
	}
	return isEqualTree(s.Left, t.Left) && isEqualTree(s.Right, t.Right)
}