package binary_tree

import "algorithm/template"

func deleteNode(root *template.TreeNode, key int) *template.TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	}

	if key == root.Val {
		left := root.Left
		right := root.Right
		if right != nil {
			// 找到 node.Left == nil 的节点
			node := right
			for node.Left != nil {
				node = node.Left
			}
			node.Left = left
			return right
		} else {
			return left
		}
	}

	return root
}
