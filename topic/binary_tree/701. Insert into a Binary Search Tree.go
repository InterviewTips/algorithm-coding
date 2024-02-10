package binary_tree

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val} // 为空则直接新建并 return
	}

	if val > root.Val {
		if root.Right == nil {
			node := &TreeNode{Val: val}
			root.Right = node
		} else {
			root.Right = insertIntoBST(root.Right, val)
		}
	}

	if val < root.Val {
		if root.Left == nil {
			node := &TreeNode{Val: val}
			root.Left = node
		} else {
			root.Left = insertIntoBST(root.Left, val)
		}
	}

	return root
}
