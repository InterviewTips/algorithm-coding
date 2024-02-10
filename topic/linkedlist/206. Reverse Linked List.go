package linkedlist

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	for head.Next != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	head.Next = pre

	return head
}
