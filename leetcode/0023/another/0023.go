package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head, node *ListNode
	if l1.Val < l2.Val {
		head = l1
		node = l1
		l1 = l1.Next
	} else {
		head = l2
		node = l2
		l2 = l2.Next
	}

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}

		node = node.Next
	}

	if l1 != nil {
		// 连上 l1 剩余的链
		node.Next = l1
	}

	if l2 != nil {
		// 连上 l2 剩余的链
		node.Next = l2
	}

	return head

}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	s := 1 // 归并步长，代表当前合并的两个链表的索引差
	for s < len(lists) {
		for i := 0; i < len(lists); i = i + 2*s {
			if i+s <= len(lists)-1 {
				lists[i] = mergeTwoLists(lists[i], lists[i+s])
			}
		}
		s = s * 2
	}

	return lists[0]

}

func main() {
	var n []*ListNode
	n = append(n, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	})
	n = append(n, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	})
	n = append(n, &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  6,
			Next: nil,
		},
	})
	r := mergeKLists(n)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
