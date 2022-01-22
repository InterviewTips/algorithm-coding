package two_pointers

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 计算 a 长度
	headALen := 0
	tmpA := headA
	for tmpA != nil {
		headALen++
		tmpA = tmpA.Next
	}
	// 计算 b 长度
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

	// 对齐
	for d != 0 {
		tmpA = tmpA.Next
		d--
	}

	for tmpA != nil {
		if tmpA == tmpB {
			return tmpA
		}

		tmpA = tmpA.Next
		tmpB = tmpB.Next
	}

	return nil

}
