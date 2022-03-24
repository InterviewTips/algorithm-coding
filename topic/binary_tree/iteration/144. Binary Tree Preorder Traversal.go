package iteration

import "algorithm/topic/binary_tree"

type TreeNode = binary_tree.TreeNode

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	// 先 push root，然后 pop root, 然后 push 右，再 push 左
	// 栈弹出的顺序为中, 左右
	stack := make([]*TreeNode, 0)
	if root == nil {
		return res
	}
	stack = append(stack, root)
	for len(stack) != 0 { // no empty
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// add val
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return res
}
