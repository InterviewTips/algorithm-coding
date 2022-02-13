package binary_tree

// type NTreeNode = Node
func preorder(root *NTreeNode) []int {
	res := make([]int, 0)
	subPreOrder(root, &res)
	return res
}

func subPreOrder(root *NTreeNode, res *[]int) {
	if root == nil {
		return
	}

	// 中左右
	*res = append(*res, root.Val)
	for i := 0; i < len(root.Children); i++ {
		subPreOrder(root.Children[i], res)
	}

	return
}

// type NTreeNode = Node
// var preorder = preorderIteration
func preorderIteration(root *NTreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	// 中左右 -> 右左(中+nil)
	stack := make([]*NTreeNode, 0)
	// push root to stack
	stack = append(stack, root)
	for len(stack) != 0 {
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			// push
			for i := len(node.Children) - 1; i >= 0; i-- {
				stack = append(stack, node.Children[i])
			}
			stack = append(stack, node)
			stack = append(stack, nil)
		} else { // nil
			// pop
			realNode := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, realNode.Val)
		}
	}

	return res
}
