package main

import (
	"fmt"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		a   string
		b   string
		c   int
		d   int
		e   int
		f   string
		j   int
		l3  *ListNode
		cur *ListNode
		err error
	)

	for ; l1 != nil; l1 = l1.Next {
		a = fmt.Sprint(l1.Val, a)
	}

	for ; l2 != nil; l2 = l2.Next {
		b = fmt.Sprint(l2.Val, b)
	}

	if c, err = strconv.Atoi(a); err != nil {
		return nil
	}

	if d, err = strconv.Atoi(b); err != nil {
		return nil
	}

	e = c + d

	f = strconv.Itoa(e)

	l3 = &ListNode{}

	cur = l3

	for i := len(f) - 1; i >= 0; i-- {
		j, err = strconv.Atoi(fmt.Sprintf("%c", f[i]))
		cur.Next = &ListNode{
			Val: j,
		}
		cur = cur.Next
	}

	return l3.Next

}

func main() {

	l1 := &ListNode{
		Val: 2,
	}
	l1.Next = &ListNode{
		Val: 4,
	}

	l1.Next.Next = &ListNode{
		Val: 3,
	}

	l2 := &ListNode{
		Val: 5,
	}

	l2.Next = &ListNode{
		Val: 6,
	}

	l2.Next.Next = &ListNode{
		Val: 4,
	}

	l3 := addTwoNumbers(l1, l2)

	for ; l3 != nil; l3 = l3.Next {
		fmt.Println(l3.Val)
	}
}
