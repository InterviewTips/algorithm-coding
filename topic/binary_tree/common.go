package binary_tree

import (
	"log"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Null 如果要表示 nil，则使用 math.MinInt64
var Null = math.MinInt64

type NTreeNode struct {
	Val      int
	Children []*NTreeNode
}

// CreateNTree 树的序列化输入是用层序遍历，每组子节点都由 null 值分隔 详情参考 unit test case
func CreateNTree(data []int) *NTreeNode {
	if len(data) == 0 {
		return nil
	}
	treeQueue := make([]*NTreeNode, 0)
	// push root to queue
	head := &NTreeNode{Val: data[0], Children: make([]*NTreeNode, 0)}
	treeQueue = append(treeQueue, head)
	index := 1
	for index < len(data) {
		// 正常来讲这里 data[index] 就一定是 null
		//if data[index] == Null { // 如果是 null 则结束此节点
		// pop_front 然后进行节点 children 的加入
		node := treeQueue[0]
		treeQueue = treeQueue[1:]
		index++
		for index < len(data) && data[index] != Null {
			tNode := &NTreeNode{Val: data[index], Children: make([]*NTreeNode, 0)}
			node.Children = append(node.Children, tNode)
			// push to queue
			treeQueue = append(treeQueue, tNode)
			index++
		}
		//} else { // != nil push node to queue
		//	tNode := &NTreeNode{Val: data[index], Children: make([]*NTreeNode, 0)}
		//	treeQueue = append(treeQueue, tNode)
		//	index++
		//}
	}

	return head
}

func LevelOrderLog(root *NTreeNode) {
	if root == nil {
		return
	}
	treeQueue := make([]*Node, 0)
	// push root to queue
	treeQueue = append(treeQueue, root)
	for len(treeQueue) != 0 {
		levelRes := make([]int, 0)
		size := len(treeQueue)
		for i := 0; i < size; i++ {
			// pop_front
			node := treeQueue[0]
			treeQueue = treeQueue[1:]
			// add children to queue
			treeQueue = append(treeQueue, node.Children...)
			// add node.Val to levelRes
			levelRes = append(levelRes, node.Val)
		}
		// append levelRes to res
		log.Println(levelRes)
	}

	return
}

//CreateBinaryTree 构造二叉树 leetcode
func CreateBinaryTree(data []int) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	q := queue{nodes: make([]*TreeNode, 0)}
	head := &TreeNode{Val: data[0]}
	q.push(head)
	index := 1 // 从下标索引 1 开始
	for index < len(data) {
		tmp := q.pop()
		if tmp == nil { // 提前返回
			return head
		}
		if data[index] != Null { // -> left
			node := &TreeNode{Val: data[index]}
			tmp.Left = node
			q.push(node) // add node
		}
		index++
		if index < len(data) && data[index] != Null { // -> right
			node := &TreeNode{Val: data[index]}
			tmp.Right = node
			q.push(node) // add node
		}
		index++
	}

	return head
}

//PreOrder 前序遍历
func PreOrder(root *TreeNode) {
	if root != nil {
		log.Println(root.Val)
		PreOrder(root.Left)
		PreOrder(root.Right)
	} else {
		return
	}
}

//PostOrder 后序遍历
func PostOrder(root *TreeNode) {
	if root != nil {
		PostOrder(root.Left)
		PostOrder(root.Right)
		log.Println(root.Val)
	} else {
		return
	}
}

//InOrder 中序遍历
func InOrder(root *TreeNode) {
	if root != nil {
		InOrder(root.Left)
		log.Println(root.Val)
		InOrder(root.Right)
	} else {
		return
	}
}

type queue struct {
	nodes []*TreeNode
}

func (q *queue) pop() *TreeNode {
	if q.empty() {
		return nil
	}
	node := q.nodes[0] // 取出首个元素
	q.nodes = q.nodes[1:]
	return node
}

func (q *queue) front() *TreeNode {
	if q.empty() {
		return nil
	}
	node := q.nodes[0] // 取出首个元素
	return node
}

func (q *queue) push(node *TreeNode) {
	q.nodes = append(q.nodes, node)
}

func (q *queue) empty() bool {
	return len(q.nodes) == 0
}
