package main

import "algorithm/guns"

type TreeNode = guns.TreeNode // 修改签名，下面的 TreeNode 就不需要变

func sumOfLeftLeaves(root *TreeNode) int {
	//return subSumOfLeftLeaves(root, false, 0)
	return subSumOfLeftLeaves1(root, false)
}

func subSumOfLeftLeaves1(node *TreeNode, left bool) (ans int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		if left {
			ans += node.Val
		}
		return
	}
	ans += subSumOfLeftLeaves1(node.Left, true)
	ans += subSumOfLeftLeaves1(node.Right, false)
	return
}

func subSumOfLeftLeaves(root *TreeNode, left bool, sum int) int {
	// 边界条件
	if root == nil {
		return sum
	}
	if root.Left == nil && root.Right == nil { //叶子节点
		if left { //累加
			sum += root.Val
		}
		return sum
	}
	sum = subSumOfLeftLeaves(root.Left, true, sum) // 这里把 sum 传进去了 所以不需要 +=
	sum = subSumOfLeftLeaves(root.Right, false, sum)
	return sum
}
