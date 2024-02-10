package binary_tree

import "algorithm/template"

func inorderTraversal(root *template.TreeNode) []int {
	res := make([]int, 0)
	subInorderTraversal(root, &res)
	return res
}

func subInorderTraversal(node *template.TreeNode, res *[]int) { // res 需要指针，不然每次地址发生变化 无法定位到
	if node == nil {
		return
	}
	// 左中右
	subInorderTraversal(node.Left, res)
	*res = append(*res, node.Val)
	subInorderTraversal(node.Right, res)

	return
}
