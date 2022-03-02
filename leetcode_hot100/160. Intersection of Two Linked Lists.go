package leetcode_hot100

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	h1 := headA
	h1Count := 0
	for h1 != nil {
		h1Count++
		h1 = h1.Next
	}
	h2 := headB
	h2Count := 0
	for h2 != nil {
		h2Count++
		h2 = h2.Next
	}

	if h1Count > h2Count {
		count := h1Count - h2Count
		// 移动 h1
		for count != 0 {
			headA = headA.Next
			count--
		}
	} else {
		count := h2Count - h1Count
		// 移动 h1
		for count != 0 {
			headB = headB.Next
			count--
		}
	}

	for headA != nil {
		if headA == headB {
			return headA
		}

		// a 和 b 都移动
		headA = headA.Next
		headB = headB.Next
	}

	return nil
}
