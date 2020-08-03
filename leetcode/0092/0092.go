package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	res := &ListNode{}
	pre := res
	pre.Next = head
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}
	var node *ListNode = nil
	cur := pre.Next
	for i := 0; i < n-m+1; i++ {
		tmp := cur.Next
		cur.Next = node
		node = cur
		cur = tmp
	}
	pre.Next.Next = cur
	pre.Next = node

	return res.Next

}

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	res := reverseBetween(node, 1, 3)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
