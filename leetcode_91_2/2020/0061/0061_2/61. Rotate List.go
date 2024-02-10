package _061

import (
	"fmt"

	"algorithm/template"
)

type ListNode = template.ListNode

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
	newTail := head
	for i := 1; i < n-k%n; i++ {
		newTail = newTail.Next
	}
	// 指向开头 顺序不能颠倒 如果颠倒， newTail 如果是最后一个节点，(其实这里可以通过判断 k % n = 0 提前返回)
	// 那么 newHead = (newTail.Next = nil)
	p.Next = head
	newHead := newTail.Next
	newTail.Next = nil
	return newHead
}
