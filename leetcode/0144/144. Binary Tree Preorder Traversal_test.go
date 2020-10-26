package _144

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_preorderTraversal(t *testing.T) {
	assert.Equal(
		t,
		[]int{1, 2, 3},
		preorderTraversal(&TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		}),
	)
}
