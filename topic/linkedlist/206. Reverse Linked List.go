package linkedlist

import "algorithm/template"

func reverseList(head *template.ListNode) *template.ListNode {
	if head == nil {
		return nil
	}
	var pre *template.ListNode
	for head.Next != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	head.Next = pre

	return head
}
