package _234

import "github.com/InterviewTips/algorithm-coding/guns"

type ListNode = guns.ListNode

func isPalindrome(head *ListNode) bool {
	//使用数组存储
	p := head
	var nodes []int
	for p != nil {
		nodes = append(nodes, p.Val)
		p = p.Next
	}
	//判断是否是回文字符串
	for i := 0; i < len(nodes)/2; i++ {
		if nodes[i] != nodes[len(nodes)-1-i] {
			return false
		}
	}
	return true
}
