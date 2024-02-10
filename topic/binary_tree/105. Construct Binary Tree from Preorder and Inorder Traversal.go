package binary_tree

import "algorithm/template"

func buildTree(preorder []int, inorder []int) *template.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &template.TreeNode{Val: preorder[0]}

	// 寻找索引
	var index int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			index = i
			break
		}
	}

	// 四个区间
	// 对于任意一颗树而言，前序遍历的形式总是
	// [ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
	// 即根节点总是前序遍历中的第一个节点。而中序遍历的形式总是
	// [ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]
	// 中序遍历的 left 和 right 好确认 :i 和 i+1:
	// 前序遍历的 left 则是 1: len(inorder[:index])+1, right 则是 index+1:
	// 在这里 len(inorder[:index]) = index

	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	//root.Left = buildTree(preorder[1:len(inorder[:index])+1], inorder[:index])
	//root.Right = buildTree(preorder[len(inorder[:index])+1:], inorder[index+1:])

	return root
}
