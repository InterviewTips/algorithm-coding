package binary_tree

import "algorithm/template"

// 自底向上的层序遍历
func levelOrderBottom(root *template.TreeNode) [][]int {
	res := levelOrder(root)
	subLevelOrderBottom(&res)
	return res
}

func subLevelOrderBottom(res *[][]int) {
	for i := 0; i < len(*res)/2; i++ {
		(*res)[i], (*res)[len(*res)-i-1] = (*res)[len(*res)-i-1], (*res)[i]
	}
}
