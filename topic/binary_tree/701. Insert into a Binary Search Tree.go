package binary_tree

import "algorithm/template"

func insertIntoBST(root *template.TreeNode, val int) *template.TreeNode {
	if root == nil {
		return &template.TreeNode{Val: val} // 为空则直接新建并 return
	}

	if val > root.Val {
		if root.Right == nil {
			node := &template.TreeNode{Val: val}
			root.Right = node
		} else {
			root.Right = insertIntoBST(root.Right, val)
		}
	}

	if val < root.Val {
		if root.Left == nil {
			node := &template.TreeNode{Val: val}
			root.Left = node
		} else {
			root.Left = insertIntoBST(root.Left, val)
		}
	}

	return root
}
