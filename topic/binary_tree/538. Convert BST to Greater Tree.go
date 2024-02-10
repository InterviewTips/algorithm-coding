package binary_tree

// todo: 其实就是需要反向中序遍历 后面 review 的时候做一下
func convertBST(root *TreeNode) *TreeNode {
	// 中序遍历 算出累加值
	inorderRes := inorderTraversal(root)
	data := make(map[int]int) // 值对应累加值
	for i := len(inorderRes) - 1; i >= 0; i-- {
		item := inorderRes[i]
		if i < len(inorderRes)-1 {
			inorderRes[i] += inorderRes[i+1]
		}
		data[item] = inorderRes[i]
	}
	//log.Println(inorderRes)
	//log.Println(data)

	subConvertBST(root, data)

	return root
}

func subConvertBST(node *TreeNode, data map[int]int) {
	if node == nil {
		return
	}
	// 左中右
	subConvertBST(node.Left, data)
	// 修改累加值
	node.Val = data[node.Val]
	subConvertBST(node.Right, data)

	return
}
