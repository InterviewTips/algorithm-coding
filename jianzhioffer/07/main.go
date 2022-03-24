package main

import "algorithm/guns"

type TreeNode = guns.TreeNode

type tree struct {
	preorder   []int
	inorder    []int
	inorderMap map[int]int
}

func newTree(preorder, inorder []int) *tree {
	return &tree{
		preorder:   preorder,
		inorder:    inorder,
		inorderMap: make(map[int]int),
	}
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// 中序遍历建立索引
	tree := newTree(preorder, inorder)
	length := len(inorder)
	for i := 0; i < length; i++ {
		tree.inorderMap[inorder[i]] = i
	}
	return tree.subBuildTree(0, length-1, 0, length-1)
}

func (t *tree) subBuildTree(preorderLeft, preorderRight,
	inorderLeft, inorderRight int) *TreeNode {
	// 边界条件
	if preorderLeft > preorderRight {
		return nil
	}
	// 找出根节点
	rootIndex := preorderLeft
	inorderRootIndex := t.inorderMap[t.preorder[rootIndex]]
	// 建议根结点
	root := &TreeNode{Val: t.preorder[rootIndex]}
	// 算出左子树有多少个元素 按照中序遍历去算
	leftTreeCount := inorderRootIndex - inorderLeft
	root.Left = t.subBuildTree(preorderLeft+1, preorderLeft+leftTreeCount,
		inorderLeft, inorderRootIndex-1)
	root.Right = t.subBuildTree(preorderLeft+leftTreeCount+1, preorderRight,
		inorderRootIndex+1, inorderRight)

	return root
}
