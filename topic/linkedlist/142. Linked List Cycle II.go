package linkedlist

import "algorithm/template"

// a + n * (b + c) + b = 2 * (a + b)
// fast 每次走两步，slow 每次走一步
func detectCycle(head *template.ListNode) *template.ListNode {
	slow := head
	fast := head
	for fast != nil {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else { // 说明没有环 直接退出
			return nil
		}
		if slow == fast { // 相遇了
			ptr := head // 从头开始 跟 slow 一起 .Next
			for ptr != slow {
				ptr = ptr.Next
				slow = slow.Next
			}
			return ptr
		}
	}

	return nil
}
