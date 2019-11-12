package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	res := &ListNode{}
	node := res
	for head != nil {
		dup := false
		for head != nil && head.Next != nil && head.Val == head.Next.Val {
			dup = true
			head = head.Next
		}
		if !dup {
			res.Next = head
			res = res.Next
		}
		head = head.Next
	}

	// 这一步需要，不然例如 1->1->2->3->3 的结果会为 2->3->3
	res.Next = nil
	return node.Next

}

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}
	res := deleteDuplicates(node)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
