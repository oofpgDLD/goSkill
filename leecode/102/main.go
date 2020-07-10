package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	s := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right:&TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right:&TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(levelOrder(s))
}

func levelOrder(root *TreeNode) [][]int {
	m := make(map[int][]int)
	BFS(m, root, 0)
	ret := make([][]int,0)
	sortedLv := make([]int,0)
	for lv := range m{
		sortedLv = append(sortedLv, lv)
	}
	sort.Ints(sortedLv)

	for _,lv := range sortedLv{
		ret = append(ret, m[lv])
	}
	return ret
}

func BFS(m map[int][]int, root *TreeNode, level int) {
	if root == nil {
		return
	}
	if store,ok := m[level];ok{
		store = append(store, root.Val)
		m[level] = store
	}else {
		store := make([]int,0)
		store = append(store, root.Val)
		m[level] = store
	}

	BFS(m,root.Left, level+1)
	BFS(m,root.Right, level+1)
}