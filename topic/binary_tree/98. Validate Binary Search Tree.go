package binary_tree

import "algorithm/template"

// 二叉搜索树 每个节点的值都不能重复
// var isValidBST = isValidBSTIteration
func isValidBSTIteration(root *template.TreeNode) bool {
	// 利用性质 中序遍历有序
	if root == nil {
		return true
	}
	stack := make([]*template.TreeNode, 0)
	res := make([]int, 0)
	// push root to stack
	stack = append(stack, root)
	for len(stack) != 0 {
		// pop node
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			stack = append(stack, node, nil)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else { // nil pop realNode
			realNode := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(res) > 0 && res[len(res)-1] >= realNode.Val { // 直接 return
				return false
			} else { // 加入 res 中
				res = append(res, realNode.Val)
			}
		}
	}

	return true
}

// 递归
func isValidBST(root *template.TreeNode) bool {
	return subIsValidBST(root, nil, nil) // 根节点没有上下界 所以传 nil
}

// 定义多一个上界和下界
// root.Left 更新上界
// root.Right 更新下界
func subIsValidBST(root *template.TreeNode, upper, lower *int) bool { // 巧妙利用 *int 指针 不然就需要定义最大值最小值之类
	if root == nil {
		return true
	}

	if lower != nil && root.Val <= *lower {
		return false
	}

	if upper != nil && root.Val >= *upper {
		return false
	}

	left := subIsValidBST(root.Left, &root.Val, lower)   // 更新上界 都要比这个上界小
	right := subIsValidBST(root.Right, upper, &root.Val) // 更新下界，都要比这个下界大

	return left && right
}
