package binary_tree

import (
	"math"
)

// 任意两个节点的最小差值，根据中序遍历来
func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}

	min := math.MaxInt64
	// 关键点 记录先前已经遍历到的最后一个节点 到时候才可以做相减
	preNode := &TreeNode{Val: -1} // 要先初始化 不能直接传 nil 多级指针修改
	subGetMinimumDifference(root, &preNode, &min)
	return min
}

// 0 <= Node.val <= 105
func subGetMinimumDifference(node *TreeNode, preNode **TreeNode, minVal *int) {
	if node == nil {
		return
	}

	// 左中右
	subGetMinimumDifference(node.Left, preNode, minVal)
	// 如果 preNode.Val == -1 说明还没有被赋值过 所以赋值一下且无需计算
	if (*preNode).Val == -1 {
		*preNode = node
	} else {
		val := node.Val - (*preNode).Val
		if val < *minVal {
			*minVal = val
		}
		*preNode = node
	}
	subGetMinimumDifference(node.Right, preNode, minVal)

	return

}
