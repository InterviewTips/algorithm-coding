package linkedlist

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	headALen := 0
	tmpA := headA
	for tmpA != nil {
		headALen++
		tmpA = tmpA.Next
	}
	headBLen := 0
	tmpB := headB
	for tmpB != nil {
		headBLen++
		tmpB = tmpB.Next
	}

	tmpA = headA
	tmpB = headB
	if headALen < headBLen {
		tmpA, tmpB = headB, headA
		headALen, headBLen = headBLen, headALen
	}

	d := headALen - headBLen

	for d != 0 {
		tmpA = tmpA.Next
		d--
	}

	for tmpA != nil { // 这里的判定条件不是 .Next != nil 有可能是最后一个节点就重合了
		if tmpA == tmpB {
			return tmpA
		}

		tmpA = tmpA.Next
		tmpB = tmpB.Next
	}

	return nil
}
