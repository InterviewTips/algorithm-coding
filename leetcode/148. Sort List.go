package leetcode

import (
	"sort"

	"algorithm/template"
)

type ListNode = template.ListNode

func sortList(head *ListNode) *ListNode {
	// 排序然后重建下

	nums := make([]int, 0)

	cur := head
	for cur != nil {
		nums = append(nums, cur.Val)
		cur = cur.Next
	}

	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	dummyHead := &ListNode{}
	pre := dummyHead
	for i := 0; i < len(nums); i++ {
		pre.Next = &ListNode{Val: nums[i]}
		pre = pre.Next
	}

	return dummyHead.Next
}
