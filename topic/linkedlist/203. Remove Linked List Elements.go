package linkedlist

// 给你一个链表的头节点 head 和一个整数 val
// 请你删除链表中所有满足 Node.val == val 的节点 并返回 新的头节点
// 输入：head = [1,2,6,3,4,5,6], val = 6
//输出：[1,2,3,4,5]
//示例 2：
//
//输入：head = [], val = 1
//输出：[]
//示例 3：
//
//输入：head = [7,7,7,7], val = 7
//输出：[]
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	cur := dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummyHead.Next
}
