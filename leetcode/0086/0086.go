package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := &ListNode{}
	temp := head
	for temp.Next != nil {
		temp = temp.Next
	}
	oldLast, last := temp, temp
	temp = res
	for head != nil {
		if head == last {
			if head.Next == nil {
				temp.Next = head
				temp = temp.Next
				head = head.Next
			} else {
				if head.Val >= x {
					temp.Next = head.Next
					//temp = temp.Next
					oldLast.Next = head
					oldLast = oldLast.Next
					oldLast.Next = nil
					head = temp.Next
				} else {
					temp.Next = head
					temp = temp.Next
					head = head.Next
				}
			}

			break
		}
		if head.Val >= x {
			temp.Next = head.Next
			//temp = temp.Next
			oldLast.Next = head
			oldLast = oldLast.Next
			oldLast.Next = nil
			head = temp.Next
		} else {
			temp.Next = head
			temp = temp.Next
			head = head.Next
		}
	}

	return res.Next

}

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  2,
							Next: nil,
						},
					},
				},
			},
		},
	}
	//node := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val:  2,
	//		Next: nil,
	//	},
	//}
	res := partition(node, 3)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
