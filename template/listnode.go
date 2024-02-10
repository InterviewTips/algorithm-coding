package template

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int, next *ListNode) *ListNode {
	return &ListNode{Val: val, Next: next}
}

func NewLinkedList(vals []int) *ListNode {
	nodes := make([]*ListNode, len(vals)+1)
	for i := len(vals) - 1; i >= 0; i-- {
		nodes[i] = NewListNode(vals[i], nodes[i+1])
	}

	return nodes[0]
}
