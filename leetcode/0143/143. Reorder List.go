package _143

import "github.com/InterviewTips/algorithm-coding/guns"

type ListNode = guns.ListNode

func reorderList(head *ListNode) {
	//边界判定
	if head == nil {
		return
	}
	//计算长度
	p := head
	nodeLens := 0
	var nodes []*ListNode
	for p != nil {
		nodes = append(nodes, p)
		nodeLens += 1
		p = p.Next
	}
	if len(nodes) == 0 {
		return
	}
	//倒序
	p = head                              // 重置
	for i := 0; i < (nodeLens-1)/2; i++ { // 边界问题，需要交换几次
		x := p.Next
		p.Next = nodes[nodeLens-(i+1)]
		nodes[nodeLens-(i+1)].Next = x
		p = x
	}
	//这里需要解决指针 .Next 问题，如果是奇数，其实 p.Next 就为 nil，不需要接后续的节点
	if nodeLens%2 == 1 { // nil
		p.Next = nil
	} else { // 偶数，说明还剩一个节点
		p.Next.Next = nil
	}
	return
}
