package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		nodeLens, count = 0, 0
		prev            *ListNode
	)
	temp := head
	for temp != nil {
		nodeLens += 1
		temp = temp.Next
	}
	fmt.Println(nodeLens)

	target := nodeLens - n
	temp = head
	for temp != nil {
		if count == target {
			if prev != nil {
				prev.Next = temp.Next
			} else {
				head = temp.Next
			}
			break
		}
		prev = temp
		temp = temp.Next
		count += 1

	}

	return head
}

func main() {
	//nodeList := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val:  3,
	//			Next: nil,
	//		},
	//	},
	//}
	//l := removeNthFromEnd(nodeList, 2)
	nodeList := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	l := removeNthFromEnd(nodeList, 2)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

}
