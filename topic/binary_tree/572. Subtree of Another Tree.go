package binary_tree

import "algorithm/template"

func isSubtree(root *template.TreeNode, subRoot *template.TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}

	mid := isSameTree(root, subRoot) // 中
	if mid {
		return true
	}
	left := isSubtree(root.Left, subRoot)
	if left {
		return true
	}
	right := isSubtree(root.Right, subRoot)
	if right {
		return true
	}

	return false
}
