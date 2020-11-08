package _109

import "github.com/InterviewTips/algorithm-coding/guns"

type (
	ListNode = guns.ListNode
	TreeNode = guns.TreeNode
)

func sortedListToBST(head *ListNode) *TreeNode {
	// 数组
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return buildBSTTree(arr, 0, len(arr)-1)
}

func buildBSTTree(arr []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	root := &TreeNode{Val: arr[mid]}
	root.Left = buildBSTTree(arr, start, mid-1)
	root.Right = buildBSTTree(arr, mid+1, end)
	return root
}
