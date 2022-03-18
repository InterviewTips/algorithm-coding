package _0220319

import (
	"fmt"

	"github.com/InterviewTips/algorithm-coding/topic/binary_tree"
)

type TreeNode = binary_tree.TreeNode

func tree2str(root *TreeNode) string {

	var res string

	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}

		res += fmt.Sprintf("%d", node.Val)
		if node.Left == nil && node.Right == nil {
			return
		}
		if node.Left != nil {
			res += "("
			preOrder(node.Left)
			res += ")"
		} else { // left 为空不能省略括号
			res += "()"
		}
		if node.Right != nil {
			res += "("
			preOrder(node.Right)
			res += ")"
		} // 右边的情况不用考虑
	}

	preOrder(root)

	return res
}
