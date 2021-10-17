package linkedlist

// 单链表和双链表的实现
type MyLinkedList struct {
	Head   *ListNode
	Length int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (l *MyLinkedList) Get(index int) int {
	if index+1 > l.Length || index < 0 { // 索引无效
		return -1
	}

	cur := &ListNode{}
	cur.Next = l.Head
	for index >= 0 {
		cur = cur.Next
		index--
	}

	return cur.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	defer func() {
		l.Length++
	}()
	a := &ListNode{Val: val, Next: l.Head}
	l.Head = a
	return
}

func (l *MyLinkedList) AddAtTail(val int) {
	defer func() {
		l.Length++
	}()
	if l.Head == nil {
		l.Head = &ListNode{Val: val}
		return
	}
	// 遍历到最后
	cur := &ListNode{}
	cur.Next = l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &ListNode{Val: val}

	return
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index > l.Length { // 大于 length 无需添加
		return
	}
	if index <= 0 {
		l.AddAtHead(val)
		return
	}
	if index == l.Length {
		l.AddAtTail(val)
		return
	}
	defer func() {
		l.Length++
	}()
	cur := &ListNode{}
	pre := &ListNode{}
	cur.Next = l.Head
	for index >= 0 {
		pre = cur
		cur = cur.Next
		index--
	}
	//pre.Next = &ListNode{Val: val, Next: cur}
	newNode := &ListNode{Val: val}
	pre.Next = newNode
	newNode.Next = cur

	return
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index+1 > l.Length || index < 0 { // 索引无效
		return
	}
	defer func() {
		l.Length--
	}()

	cur := &ListNode{}
	pre := &ListNode{}
	cur.Next = l.Head
	dummy := cur
	for index >= 0 {
		pre = cur
		cur = cur.Next
		index--
	}
	pre.Next = cur.Next
	l.Head = dummy.Next

	return
}
