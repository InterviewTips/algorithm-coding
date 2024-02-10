package template

import "testing"

func TestGenLinkList(t *testing.T) {
	head := GenLinkList([]int{1, 2, 1})
	for head != nil {
		t.Logf("head is %v\n", head)
		head = head.Next
	}
}
