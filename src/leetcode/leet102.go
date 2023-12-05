package leetcode

import "fmt"

func listLeft(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	r := []int{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		r = append(r, q[0].Val)
		q2 := []*TreeNode{}
		for _, n := range q {
			if n.Left != nil {
				q2 = append(q2, n.Left)
			}
			if n.Right != nil {
				q2 = append(q2, n.Right)
			}
		}
		q = q2
	}
	return r
}

func test102() {
	rt := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}}
	fmt.Print(listLeft(rt))
	fmt.Println()
}
