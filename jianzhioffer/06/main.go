package main

import (
	"github.com/InterviewTips/algorithm-coding/guns"
)

type ListNode = guns.ListNode

func reversePrint(head *ListNode) []int {
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	i := 0
	j := len(res) - 1
	for i <= j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res

}
