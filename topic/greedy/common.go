package greedy

import (
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

// CreateBinaryTree 构造二叉树 leetcode
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
		if index < len(data) && data[index] != Null { // -> left
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
