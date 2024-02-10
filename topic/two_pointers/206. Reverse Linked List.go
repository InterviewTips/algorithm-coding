package two_pointers

import (
	"algorithm/template"
)

type ListNode = template.ListNode

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	var next *ListNode

	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}

	return pre
}
