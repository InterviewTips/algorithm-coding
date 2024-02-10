package binary_tree

func searchBST(root *TreeNode, val int) *TreeNode {
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
