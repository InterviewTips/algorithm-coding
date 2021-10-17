package linkedlist

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1, 2)          //链表变为1-> 2-> 3
	t.Log("get 1 is", linkedList.Get(1)) //返回2
	linkedList.DeleteAtIndex(1)          //现在链表是1-> 3
	t.Log("get 1 is", linkedList.Get(1)) //返回3
}

func TestCase1(t *testing.T) {
	linkedList := Constructor()
	linkedList.AddAtHead(7)
	linkedList.AddAtHead(2)
	linkedList.AddAtHead(1)
	linkedList.AddAtIndex(3, 0)
	linkedList.DeleteAtIndex(2)
	linkedList.AddAtHead(6)
	linkedList.AddAtTail(4)
	t.Log("get 4 is", linkedList.Get(4))
	linkedList.AddAtHead(4)
	linkedList.AddAtIndex(5, 0)
	linkedList.AddAtHead(6)
}

func TestCase2(t *testing.T) {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1, 2)
	t.Log(linkedList.Get(1))
	linkedList.DeleteAtIndex(0)
	t.Log(linkedList.Get(0))
}
