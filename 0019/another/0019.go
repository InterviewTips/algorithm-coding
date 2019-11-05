package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	prev := &ListNode{}
	prev.Next = head
	var start, end = prev, prev
	for n != 0 {
		start = start.Next
		n--
	}
	for start.Next != nil {
		start = start.Next
		end = end.Next
	}
	end.Next = end.Next.Next

	return prev.Next
}

func main() {
	nodeList := &ListNode{
		Val: 1,
		Next: nil,
	}
	l := removeNthFromEnd(nodeList, 1)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
