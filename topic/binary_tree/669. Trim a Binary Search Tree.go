package binary_tree

import "algorithm/template"

func trimBST(root *template.TreeNode, low int, high int) *template.TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < low { // root 左子树已经已经不符合需求
		right := trimBST(root.Right, low, high)
		return right
	}

	if root.Val > high {
		left := trimBST(root.Left, low, high)
		return left
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)

	return root
}
