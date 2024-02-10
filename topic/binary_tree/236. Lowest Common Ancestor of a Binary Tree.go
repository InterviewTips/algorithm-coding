package binary_tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}

	// 后序遍历 左右中
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if (left == q && right == p) || (left == p && right == q) {
		return root
	}

	if left == nil && right == nil {
		return nil
	}

	if left == nil { // right 不为空
		return right
	}

	if right == nil { // left 不为空
		return left
	}

	// 理论上走不到这里
	return nil
}
