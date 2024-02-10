package binary_tree

import "algorithm/template"

func searchBST(root *template.TreeNode, val int) *template.TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if val < root.Val {
		return searchBST(root.Left, val)
	} else { // if val > root.Val
		return searchBST(root.Right, val)
	}

}
