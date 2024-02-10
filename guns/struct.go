package guns

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//TODO: gen tree

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// GenLinkList 生成链表
func GenLinkList(nodes []int) *ListNode {
	p := &ListNode{}
	q := p
	for i := 0; i < len(nodes); i++ {
		q.Next = &ListNode{
			Val:  nodes[i],
			Next: nil,
		}
		q = q.Next
	}
	return p.Next
}
