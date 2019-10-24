package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resPre := &ListNode{}
	cur := resPre
	// 进位 默认为 0
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return resPre.Next
}

func main() {
	l1 := &ListNode{
		Val: 2,
	}
	l1.Next = &ListNode{
		Val: 4,
	}

	l1.Next.Next = &ListNode{
		Val: 3,
	}

	l2 := &ListNode{
		Val: 5,
	}

	l2.Next = &ListNode{
		Val: 6,
	}

	l2.Next.Next = &ListNode{
		Val: 4,
	}

	l3 := addTwoNumbers(l1, l2)

	for ; l3 != nil; l3 = l3.Next {
		fmt.Println(l3.Val)
	}
}
