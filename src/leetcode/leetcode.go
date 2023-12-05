// 力扣官网：https://leetcode.cn
package leetcode

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 创建二叉树
func createTree(a []int, i int) *TreeNode {
	if i < 0 || i >= len(a) || a[i] <= 0 {
		return nil // 空节点
	}
	// fmt.Println(i, a[i])
	node := &TreeNode{a[i], nil, nil}
	node.Left = createTree(a, i*2+1)
	node.Right = createTree(a, i*2+2)
	return node
}

// 前序遍历二叉树
func preOrder(root *TreeNode) {
	if root != nil {
		fmt.Print(root.Val, " ")
		preOrder(root.Left)
		preOrder(root.Right)
	}
}

// 中序遍历二叉树
func inOrder(root *TreeNode) {
	if root != nil {
		inOrder(root.Left)
		fmt.Print(root.Val, " ")
		inOrder(root.Right)
	}
}

// 后序遍历二叉树
func postOrder(root *TreeNode) {
	if root != nil {
		postOrder(root.Left)
		postOrder(root.Right)
		fmt.Print(root.Val, " ")
	}
}

// 深度优先搜索，遍历二叉树
func dfsTrans(root *TreeNode) {
	fmt.Print("先序遍历: ")
	preOrder(root)
	fmt.Println()
	fmt.Print("中序遍历: ")
	inOrder(root)
	fmt.Println()
	fmt.Print("后序遍历: ")
	postOrder(root)
	fmt.Println()
}

func Demo() {
	test102()
	test637()
}
