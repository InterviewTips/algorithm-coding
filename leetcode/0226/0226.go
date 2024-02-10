package main

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	//l := root.Left // 暂存 因为下一步被赋值
	//root.Left = invertTree(root.Right)
	//root.Right = invertTree(l)
	r := root.Right
	root.Right = invertTree(root.Left)
	root.Left = invertTree(r)

	return root
}

func main() {

}
