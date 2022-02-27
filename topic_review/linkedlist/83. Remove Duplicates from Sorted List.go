package linkedlist

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Val: -101}
	dummyHead.Next = head
	pre := dummyHead
	for pre.Next != nil {
		if pre.Next.Val != pre.Val { // 不等于则可以移动到下一个节点
			pre = pre.Next
		} else { // 如果 pre.Next 等于前面的值，则说明要指向后一个节点接着比较
			pre.Next = pre.Next.Next
		}
	}

	return dummyHead.Next
}
