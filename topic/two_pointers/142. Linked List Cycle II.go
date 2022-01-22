package two_pointers

// a + n * (b + c) + b = 2 * (a + b)
// fast 每次走两步，slow 每次走一步
func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil {
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return nil
		}
		if slow == fast {
			ptr := head
			for ptr != slow {
				slow = slow.Next
				ptr = ptr.Next
			}
			return ptr
		}
	}

	return nil
}
