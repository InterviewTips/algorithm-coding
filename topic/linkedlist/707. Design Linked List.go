package linkedlist

type MyLinkedList struct {
	DummyHead *ListNode
	Length    int
}

func Constructor() MyLinkedList {
	return MyLinkedList{DummyHead: &ListNode{}} // 初始化伪节点
}

func (l *MyLinkedList) Get(index int) int {
	if index+1 > l.Length || index < 0 { // 索引无效
		return -1
	}

	cur := l.DummyHead
	for index >= 0 {
		cur = cur.Next
		index--
	}

	return cur.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
	return
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.Length, val)
	return
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index > l.Length { // 大于 length 无需添加
		return
	}
	defer func() {
		l.Length++
	}()
	if index < 0 {
		index = 0
	}
	pre := l.DummyHead
	for index > 0 {
		pre = pre.Next
		index--
	}
	newNode := &ListNode{Val: val}
	newNode.Next = pre.Next
	pre.Next = newNode

	return
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index+1 > l.Length || index < 0 { // 索引无效
		return
	}
	defer func() {
		l.Length--
	}()

	cur := l.DummyHead
	for index > 0 {
		index--
		cur = cur.Next
	}
	cur.Next = cur.Next.Next

	return
}
