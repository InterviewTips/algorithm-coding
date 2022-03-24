package dp

import (
	"algorithm/topic/binary_tree"
)

type TreeNode = binary_tree.TreeNode

// var rob = rob3
func rob3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	val1 := root.Val
	if root.Left != nil {
		val1 += rob3(root.Left.Left) + rob3(root.Left.Right)
	}
	if root.Right != nil {
		val1 += rob3(root.Right.Left) + rob3(root.Right.Right)
	}

	val2 := rob3(root.Left) + rob3(root.Right)

	return getMax(val1, val2)
}

// var rob = memoRob
func memoRob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	memo := make(map[*TreeNode]int)

	subRob3(root, memo)

	return memo[root]
}

func subRob3(root *TreeNode, memo map[*TreeNode]int) int {
	v, ok := memo[root]
	if ok {
		return v
	}

	if root == nil {
		memo[root] = 0
		return 0
	}
	if root.Left == nil && root.Right == nil {
		memo[root] = root.Val
		return root.Val
	}

	val1 := root.Val
	if root.Left != nil {
		val1 += subRob3(root.Left.Left, memo) + subRob3(root.Left.Right, memo)
	}
	if root.Right != nil {
		val1 += subRob3(root.Right.Left, memo) + subRob3(root.Right.Right, memo)
	}

	val2 := subRob3(root.Left, memo) + subRob3(root.Right, memo)

	value := getMax(val1, val2)
	memo[root] = value
	return value
}

// var rob = dpRob
func dpRob(root *TreeNode) int {
	res := subDpRob(root)
	return getMax(res[0], res[1])
}

func subDpRob(root *TreeNode) [2]int {
	if root == nil {
		return [2]int{0, 0}
	}

	// 后序遍历
	left := subDpRob(root.Left)
	right := subDpRob(root.Right)

	val1 := root.Val + left[0] + right[0]

	val2 := getMax(left[0], left[1]) + getMax(right[0], right[1])

	return [2]int{val2, val1}
}
