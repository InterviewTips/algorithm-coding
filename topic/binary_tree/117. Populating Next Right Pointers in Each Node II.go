package binary_tree

// 方法名修改 connect2 -> connect
// var connect = connect2
// type PerfectNode = Node
func connect2(root *PerfectNode) *PerfectNode {
	if root == nil {
		return nil
	}
	cur := root
	for cur != nil {
		dummyHead := &PerfectNode{}
		pre := dummyHead
		for cur != nil {
			if cur.Left != nil {
				pre.Next = cur.Left
				pre = pre.Next
			}
			if cur.Right != nil {
				pre.Next = cur.Right
				pre = pre.Next
			}
			cur = cur.Next // 移动到当前行的下一个节点
		}
		cur = dummyHead.Next // 移动到下一行的头个节点
	}

	return root
}
