package linkedlist

// 这道题目有点奇怪，没有说明顺序一致即可，其实并没有真正删除对应节点 而是删除了后一个
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
