package main

import (
	"algorithm/template"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = template.ListNode

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		cur = next
		prev = cur // 位移
	}

	return prev
}

//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//
//	// 双指针
//	pre := &ListNode{Val: head.Val, Next: nil}
//	for head != nil {
//		cur := head
//		head = head.Next
//		pre = head
//		pre.Next = cur
//	}
//	return pre
//
//}
