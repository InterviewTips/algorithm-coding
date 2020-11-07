package _024

import (
	"fmt"

	"github.com/InterviewTips/algorithm-coding/guns"
)

type ListNode = guns.ListNode

func swapPairs(head *ListNode) *ListNode {
	last := &ListNode{}
	last.Next = head
	prev := last
	for prev.Next != nil && prev.Next.Next != nil {
		fmt.Printf("head is %v\n", head)
		start := prev.Next
		end := prev.Next.Next
		prev.Next = end
		start.Next = end.Next
		end.Next = start
		prev = start
	}
	return last.Next
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}
