package two_pointers

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else { // nil 直接退出
			return false
		}
		if slow == fast { // 相遇了
			return true
		}
	}

	return false
}
