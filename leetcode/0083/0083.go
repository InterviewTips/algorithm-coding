package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp := head
	left, right := temp, temp.Next
	for right != nil {
		if right.Val == left.Val {
			left.Next = right.Next
			right = right.Next
		} else {
			left = left.Next
			right = right.Next
		}
	}

	return head
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
