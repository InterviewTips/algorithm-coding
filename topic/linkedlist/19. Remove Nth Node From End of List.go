package linkedlist

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
// 进阶：你能尝试使用一趟扫描实现吗？
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	slow := dummyHead
	fast := head // 这里指向 head fast 就只需要走 n-1 步，如果指向 dummyHead 就需要走 n 步
	for i := 0; i < n-1; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummyHead.Next
}
