package _0220220

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummyHead := &ListNode{}
	dummyHead.Next = head
	pre := dummyHead

	count := 0
	for pre.Next != nil {
		if pre.Next.Val == 0 {
			if count == 0 {
				pre.Next = pre.Next.Next
			} else {
				pre.Next.Val = count
				pre = pre.Next // 移动到下一个节点
				count = 0      // 重置
			}
		} else {
			count += pre.Next.Val
			pre.Next = pre.Next.Next
		}
	}

	return dummyHead.Next
}
