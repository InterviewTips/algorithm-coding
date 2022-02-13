package binary_tree

func postorder(root *NTreeNode) []int {
	res := make([]int, 0)
	subPostorder(root, &res)
	return res
}

func subPostorder(root *NTreeNode, res *[]int) {
	if root == nil {
		return
	}

	// 左右中
	for i := 0; i < len(root.Children); i++ {
		subPostorder(root.Children[i], res)
	}
	*res = append(*res, root.Val)

	return
}

// type NTreeNode = Node
// var postorder = postorderIteration
func postorderIteration(root *NTreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	// 左右中 -> (中+nil)右左
	stack := make([]*NTreeNode, 0)
	// push root to stack
	stack = append(stack, root)
	for len(stack) != 0 {
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			// push
			stack = append(stack, node)
			stack = append(stack, nil)
			for i := len(node.Children) - 1; i >= 0; i-- {
				stack = append(stack, node.Children[i])
			}
		} else { // nil
			// pop
			realNode := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, realNode.Val)
		}
	}

	return res
}
