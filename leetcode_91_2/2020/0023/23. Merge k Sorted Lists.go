package _023

import (
	"sort"

	"algorithm/template"
)

type ListNode = template.ListNode

type Queue struct {
	element []*ListNode
}

func NewQueue() *Queue {
	return &Queue{element: make([]*ListNode, 0)}
}

func (q *Queue) Add(node *ListNode) {
	q.element = append(q.element, node)
	sort.SliceStable(q.element, func(i, j int) bool {
		return q.element[i].Val > q.element[j].Val // 最小的放最后
	})
}

func (q *Queue) Pop() *ListNode {
	e := q.element[q.len()-1]
	q.element = q.element[:q.len()-1]
	return e
}

func (q *Queue) len() int {
	return len(q.element)
}

func (q *Queue) isEmpty() bool {
	return len(q.element) == 0
}

func mergeKLists(lists []*ListNode) *ListNode {
	queue := NewQueue()
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			queue.Add(lists[i])
		}
	}
	preHead := &ListNode{}
	temp := preHead
	for !queue.isEmpty() {
		node := queue.Pop()
		temp.Next = node
		temp = temp.Next
		if node.Next != nil { // node.Next 加入优先队列
			queue.Add(node.Next)
		}
	}
	temp.Next = nil // 这里其实肯定是 nil
	return preHead.Next
}
