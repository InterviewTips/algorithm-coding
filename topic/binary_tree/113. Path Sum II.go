package binary_tree

import "algorithm/template"

func pathSum(root *template.TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	path := make([]int, 0)
	path = append(path, root.Val)
	subPathSum(root, &path, &res, targetSum-root.Val)
	return res
}

func subPathSum(node *template.TreeNode, path *[]int, res *[][]int, value int) {
	if node.Left == nil && node.Right == nil && value == 0 {
		// add path to res
		// path 的对应地址的内容会被修改 所以要重新引出一个
		// 而且不是 newPath = *path， 这样仍然会被改动到 []int
		newPath := make([]int, len(*path))
		for i := 0; i < len(*path); i++ {
			newPath[i] = (*path)[i]
		}
		*res = append(*res, newPath)
		return
	}

	if node.Left != nil {
		*path = append(*path, node.Left.Val)
		subPathSum(node.Left, path, res, value-node.Left.Val)
		*path = (*path)[:len(*path)-1]
	}

	if node.Right != nil {
		*path = append(*path, node.Right.Val)
		subPathSum(node.Right, path, res, value-node.Right.Val)
		*path = (*path)[:len(*path)-1]
	}

}
