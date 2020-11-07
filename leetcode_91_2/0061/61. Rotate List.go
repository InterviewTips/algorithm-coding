package _061

import (
	"fmt"

	"github.com/InterviewTips/algorithm-coding/guns"
)

type ListNode = guns.ListNode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	n := 1
	for p.Next != nil {
		n++
		p = p.Next
	}
	fmt.Printf("len(listNode) is %v\n", n)
	p.Next = head //串成环
	// 新的链表头 n - k
	// 新的链表尾 n - k - 1
	newTail := head
	for i := 0; i < n-k%n-1; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil
	return newHead
}
