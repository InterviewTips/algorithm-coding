package binary_tree

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < low || root.Val > high {
		left := root.Left
		right := root.Right
		if right != nil {
			// 找到 node.Left == nil 的节点
			node := right
			for node.Left != nil {
				node = node.Left
			}
			node.Left = left
			// 这里不能直接 return right 需要将 right 再剪枝
			return trimBST(right, low, high)
		} else {
			// 这里不能直接 return left 需要将 left 再剪枝
			return trimBST(left, low, high)
		}
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)

	return root
}
