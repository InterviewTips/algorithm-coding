package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestSumOfLeftLeaves(t *testing.T) {
	t.Run("two left leaves", func(t *testing.T) {
		assert.Equal(t, 24, sumOfLeftLeaves(&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		}),
			"must be 24",
		)
	})
	t.Run("nil", func(t *testing.T) {
		assert.Equal(t, 0, sumOfLeftLeaves(nil), "must be 0")
	})
	// 只有一个元素的话，也是返回 0 不算是子节点
	t.Run("only root", func(t *testing.T) {
		assert.Equal(t, 0, sumOfLeftLeaves(&TreeNode{
			Val: 1,
		}), "must be 0")
	})
	t.Run("1,2", func(t *testing.T) {
		assert.Equal(t, 2, sumOfLeftLeaves(&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		}), "must be 2")
	})
}
