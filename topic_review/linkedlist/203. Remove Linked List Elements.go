package linkedlist

import (
	"algorithm/topic/linkedlist"
)

type ListNode = linkedlist.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	// 如果声明了 dummyHead 则应该修改的就是 dummyHead 而不是 head
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	for pre.Next != nil { // pre.Next == nil 说明到达末尾 此时 pre 指向最后一个节点
		if pre.Next.Val == val { // 跳过节点
			pre.Next = pre.Next.Next
		} else { // 移动到下一个节点
			pre = pre.Next
		}
	}

	return dummy.Next
}
