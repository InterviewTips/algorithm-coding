package _142

import "algorithm/guns"

type ListNode = guns.ListNode

func detectCycle1(head *ListNode) *ListNode {
	arr := make(map[*ListNode]int)
	for head != nil {
		_, ok := arr[head]
		if ok {
			return head
		}
		arr[head] = 1
		head = head.Next
	}
	return nil
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for slow != p {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
