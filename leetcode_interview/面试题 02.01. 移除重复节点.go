package leetcode_interview

import (
	"algorithm/template"
)

type ListNode = template.ListNode

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	data := make(map[int]struct{})
	data[head.Val] = struct{}{}
	pos := head
	for pos.Next != nil {
		cur := pos.Next
		_, ok := data[cur.Val]
		if !ok {
			pos = pos.Next
			data[cur.Val] = struct{}{}
			continue
		}
		// exists
		pos.Next = pos.Next.Next
	}

	return head
}
