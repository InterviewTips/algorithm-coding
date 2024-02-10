package binary_tree

import "algorithm/template"

// 提交时改名 buildTreePostOrder -> buildTree
// var buildTree = buildTreePostOrder
func buildTreePostOrder(inorder []int, postorder []int) *template.TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	root := &template.TreeNode{Val: postorder[len(postorder)-1]}

	// find index
	var index int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			index = i
			break
		}
	}

	// 中序 左 中 右
	// 后序 左 右 中
	// 正推
	root.Left = buildTreePostOrder(inorder[:index], postorder[:index])
	root.Right = buildTreePostOrder(inorder[index+1:], postorder[index:len(postorder)-1])
	// 从后面逆推
	//root.Left = buildTreePostOrder(inorder[:index], postorder[:len(postorder)-1-len(inorder[index+1:])])
	//root.Right = buildTreePostOrder(inorder[index+1:], postorder[len(postorder)-1-len(inorder[index+1:]):len(postorder)-1])

	return root
}
