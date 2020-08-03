package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	p := &ListNode{}
	p.Next = head
	var a, b, tmp = p, p, p
	// 利用 && 短路判断
	// for b.Next.Next != nil && b.Next != nil && b != nil { 这样写会报错
	// panic: runtime error: invalid memory address or nil pointer dereference
	for b != nil && b.Next != nil && b.Next.Next != nil {
		a = a.Next
		b = b.Next.Next

		tmp.Next = b
		a.Next = b.Next
		b.Next = a

		tmp = a
		b = a
	}

	return p.Next

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
	n := swapPairs(node)
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}
