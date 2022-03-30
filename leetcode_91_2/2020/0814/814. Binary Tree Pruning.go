package _814

import "algorithm/guns"

type TreeNode = guns.TreeNode

func pruneTree(root *TreeNode) *TreeNode {
	if containsOne(root) {
		return root
	} else {
		return nil
	}
}

//containsOne 子树是否包含 1
func containsOne(root *TreeNode) bool {
	if root == nil {
		return false
	}
	left := containsOne(root.Left)
	right := containsOne(root.Right)
	if !left { // 不包含
		root.Left = nil
	}
	if !right { // 不包含
		root.Right = nil
	}
	return root.Val == 1 || left || right
}
