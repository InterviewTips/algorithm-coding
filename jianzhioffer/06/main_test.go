package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReversePrint(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2}}}
	assert.Equal(t, []int{2, 3, 1}, reversePrint(head))
}
