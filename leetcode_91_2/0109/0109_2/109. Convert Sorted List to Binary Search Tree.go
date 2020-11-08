package _109

import "github.com/InterviewTips/algorithm-coding/guns"

type (
	ListNode = guns.ListNode
	TreeNode = guns.TreeNode
)

var h *ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	h = head
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return buildBSTTree(0, length-1)
}

func buildBSTTree(start, end int) *TreeNode {
	// 左中右
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	left := buildBSTTree(start, mid-1) // 左
	root := &TreeNode{Val: h.Val}      // 中
	h = h.Next
	root.Left = left

	root.Right = buildBSTTree(mid+1, end) // 右
	return root
}
