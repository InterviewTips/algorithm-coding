package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(a, b *ListNode) *ListNode {
	var pre, cur, nxt *ListNode = nil, a, a
	for cur != b {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	var a, b = head, head

	for i := 0; i < k; i++ {
		// 不足 k 个，不需要反转
		if b == nil {
			return head
		}
		b = b.Next
	}

	newHead := reverse(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead

}

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	r := reverseKGroup(node, 3)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
