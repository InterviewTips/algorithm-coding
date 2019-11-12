package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	before := &ListNode{}
	less := before
	after := &ListNode{}
	great := after
	for head != nil {
		if head.Val < x {
			less.Next = head
			less = less.Next
		} else {
			great.Next = head
			great = great.Next
		}
		head = head.Next
	}
	great.Next = nil // 避免环形链表
	less.Next = after.Next

	return before.Next

}

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  2,
							Next: nil,
						},
					},
				},
			},
		},
	}
	res := partition(node, 3)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
