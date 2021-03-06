package sample01

import "github.com/InterviewTips/algorithm-coding/guns"

type Value struct {
	maxLevel int
	res      int
}

type TreeNode = guns.TreeNode

func findBottomLeftValue(root *TreeNode) int {
	v := Value{
		maxLevel: 0,
		res:      -1,
	}
	sub(root, 1, &v)
	return v.res
}

func sub(root *TreeNode, level int, value *Value) {
	if root == nil {
		return
	}
	val := root.Val
	if level > value.maxLevel {
		value.maxLevel = level
		value.res = val
	}
	sub(root.Left, level+1, value)
	sub(root.Right, level+1, value)
}
