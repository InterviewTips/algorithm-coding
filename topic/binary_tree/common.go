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

// 如果要表示 nil，则使用 math.MinInt64
var null = math.MinInt64

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
		if data[index] != null { // -> left
			node := &TreeNode{Val: data[index]}
			tmp.Left = node
			q.push(node) // add node
		}
		index++
		if index < len(data) && data[index] != null { // -> right
			node := &TreeNode{Val: data[index]}
			tmp.Right = node
			q.push(node) // add node
		}
		index++
	}

	return head
}

//LogTree 中序遍历
func LogTree(root *TreeNode) {
	if root != nil {
		LogTree(root.Left)
		log.Println(root.Val)
		LogTree(root.Right)
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
	node := q.nodes[len(q.nodes)-1]
	q.nodes = q.nodes[:len(q.nodes)-1]
	return node
}

func (q *queue) front() *TreeNode {
	if q.empty() {
		return nil
	}
	node := q.nodes[len(q.nodes)-1]
	return node
}

func (q *queue) push(node *TreeNode) {
	q.nodes = append(q.nodes, node)
}

func (q *queue) empty() bool {
	return len(q.nodes) == 0
}
