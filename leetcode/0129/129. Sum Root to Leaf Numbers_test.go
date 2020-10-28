package _129

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_sumNumbers(t *testing.T) {
	assert.Equal(t, 25, sumNumbers(&TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}))
}
