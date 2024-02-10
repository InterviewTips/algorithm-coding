package binary_tree

import "algorithm/template"

func hasPathSum(root *template.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return subHasPathSum(root, targetSum-root.Val)
}

func subHasPathSum(node *template.TreeNode, value int) bool {
	// 边界条件 遇到叶子节点 return
	if node.Left == nil && node.Right == nil && value == 0 {
		return true
	}
	if node.Left == nil && node.Right == nil {
		return false
	}

	if node.Left != nil {
		if subHasPathSum(node.Left, value-node.Left.Val) {
			return true
		}
	}

	if node.Right != nil {
		if subHasPathSum(node.Right, value-node.Right.Val) {
			return true
		}
	}

	// 如果左右子树都不为空且不为 true 那么走到这里返回 false
	return false
}
