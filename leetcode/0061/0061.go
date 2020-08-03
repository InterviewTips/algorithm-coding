package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var n = 1
	temp := head
	for temp.Next != nil {
		n++
		temp = temp.Next
	}
	temp.Next = head

	newTail := head
	// step 为新的链表尾 n - (k % n) - 1 链表首为 n - (k % n)
	step := n - (k % n) - 1
	for i := 0; i < step; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil

	return newHead

}

func main() {
	l1 := &ListNode{
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
	node := rotateRight(l1, 2)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}

}
