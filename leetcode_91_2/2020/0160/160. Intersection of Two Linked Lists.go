package _160

import (
	"fmt"

	"algorithm/template"
)

type ListNode = template.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := 0
	tmpA := headA
	for tmpA != nil {
		lenA++
		tmpA = tmpA.Next
	}
	lenB := 0
	tmpB := headB
	for tmpB != nil {
		lenB++
		tmpB = tmpB.Next
	}
	fmt.Printf("lenA is %v, lenB is %v\n", lenA, lenB)
	if lenA > lenB { // A 先走
		diff := lenA - lenB
		for diff > 0 {
			headA = headA.Next
			diff--
		}
	} else if lenA < lenB { // B 先走
		diff := lenB - lenA
		for diff > 0 {
			headB = headB.Next
			diff--
		}
	} else { // lenA == lenB
	}
	for headA != nil {
		// 注：这里的判断条件十分重要，是 headA == headB
		// 而不是headA.Next == headB.Next && headA.Val == headB.Val 两个不同的节点也可能满足这个条件，例如示范 1 中的两个元素 1
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}
