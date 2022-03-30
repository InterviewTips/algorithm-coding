package _455

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindContentChildren(t *testing.T) {
	assert.Equal(
		t,
		1,
		findContentChildren([]int{1, 2, 3}, []int{1, 1}),
		"must be 1",
	)
	assert.Equal(
		t,
		2,
		findContentChildren([]int{1, 2}, []int{1, 2, 3}),
		"must be 1",
	)
}
