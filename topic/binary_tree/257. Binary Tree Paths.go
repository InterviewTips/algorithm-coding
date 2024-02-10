package binary_tree

import (
	"bytes"
	"fmt"
)

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	if root == nil {
		return res
	}

	path := make([]int, 0)
	subBinaryTreePaths(root, &path, &res)

	return res
}

// node 必不可能是 nil 条件已经限制
func subBinaryTreePaths(node *TreeNode, path *[]int, res *[]string) {
	// 先 push 到 path 先
	*path = append(*path, node.Val)
	// 判断是否叶子节点
	if node.Left == nil && node.Right == nil {
		// 加入 res 形式: 1->2
		var pathStr bytes.Buffer
		for i := 0; i < len(*path)-1; i++ {
			pathStr.WriteString(fmt.Sprintf("%d->", (*path)[i]))
		}
		pathStr.WriteString(fmt.Sprintf("%d", (*path)[len(*path)-1]))
		*res = append(*res, pathStr.String())
	}

	// 左节点不为空
	if node.Left != nil {
		subBinaryTreePaths(node.Left, path, res)
		*path = (*path)[:len(*path)-1] // 回溯 其实就是回退回来
	}

	// 右节点不为空
	if node.Right != nil {
		subBinaryTreePaths(node.Right, path, res)
		*path = (*path)[:len(*path)-1] // 回溯 其实就是回退回来
	}
}
