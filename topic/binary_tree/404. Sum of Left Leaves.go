package binary_tree

import "algorithm/template"

func sumOfLeftLeaves(root *template.TreeNode) int {
	sum := 0
	if root == nil {
		return sum
	}
	subSumOfLeftLeaves(root, &sum)
	return sum
}

// 传入的 root 不能为空
func subSumOfLeftLeaves(root *template.TreeNode, sum *int) {
	// 左节点不为空 且左节点是叶子节点
	if root.Left != nil && root.Left.Left == nil &&
		root.Left.Right == nil {
		*sum += root.Left.Val // 注意是加叶子节点的 value
		// return 不需要 return 可能还有别的
	}
	// 其实如果上面的条件成立，这里就没有必要往下走了
	if root.Left != nil {
		subSumOfLeftLeaves(root.Left, sum)
	}
	if root.Right != nil {
		subSumOfLeftLeaves(root.Right, sum)
	}

	return

}
