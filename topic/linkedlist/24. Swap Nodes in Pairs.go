package linkedlist

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	dummyHead.Next = head
	for cur.Next != nil && cur.Next.Next != nil { // 短路原则 前一个不成立会直接返回
		curNext := cur.Next
		curNextNext := cur.Next.Next
		curNextNextNext := cur.Next.Next.Next // 可以为 nil
		cur.Next = curNextNext
		curNextNext.Next = curNext
		curNext.Next = curNextNextNext
		cur = curNext
	}

	return dummyHead.Next
}
