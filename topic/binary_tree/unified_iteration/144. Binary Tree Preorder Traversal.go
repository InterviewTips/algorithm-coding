package iteration

import "github.com/InterviewTips/algorithm-coding/topic/binary_tree"

type TreeNode = binary_tree.TreeNode

// 中左右 -> 压入栈中的顺序是 右左(中+nil)
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil { // push to stack
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			stack = append(stack, node)
			stack = append(stack, nil)
		} else { // push to res
			// pop real node
			realNode := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, realNode.Val)
		}
	}

	return res
}
