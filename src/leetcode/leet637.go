package leetcode

import "fmt"

// 层次遍历二叉树
func averageOfLevels(root *TreeNode) []float64 {
	r := []float64{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		t := 0.0 // total
		q2 := []*TreeNode{}
		for _, n := range q {
			t += float64(n.Val)
			if n.Left != nil {
				q2 = append(q2, n.Left)
			}
			if n.Right != nil {
				q2 = append(q2, n.Right)
			}
		}
		t = t / float64(len(q))
		r = append(r, t)
		q = q2
	}
	return r
}

func test637() {
	// fmt.Println("hello world")
	if false {
		// rt := &TreeNode{3, &TreeNode{9, &TreeNode{15, nil, nil},
		// 	&TreeNode{7, nil, nil}}, &TreeNode{20, nil, nil}}
		a := []int{3, 9, 20, 15, 7}
		rt := createTree(a, 0)
		// dfsTrans(rt)
		// [3.00000,14.50000,11.00000]
		r := averageOfLevels(rt)
		for _, v := range r {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
	if true {
		// rt := &TreeNode{3, &TreeNode{9, nil, nil},
		// 	&TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
		a := []int{3, 9, 20, 0, 0, 15, 7}
		rt := createTree(a, 0)
		// dfsTrans(rt)
		// [3.00000,14.50000,11.00000]
		r := averageOfLevels(rt)
		for _, v := range r {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}
