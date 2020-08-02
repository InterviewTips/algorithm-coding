package _101

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//[1,2,2,3,4,4,3]
//[1,2,2,null,3,null,3]
//[]
func TestIsSymmetric(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		r := isSymmetric(nil)
		assert.Equal(t, true, r, "nil 为 true")
	})
	t.Run("true", func(t *testing.T) {
		tree := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		}
		r := isSymmetric(tree)
		assert.Equal(t, true, r, "对称树为 true")
	})
	t.Run("false", func(t *testing.T) {
		tree := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		}
		r := isSymmetric(tree)
		assert.Equal(t, false, r, "非对称树为 false")
	})
}

//TODO:构建树
func BuildTree(nums []int) *TreeNode {
	return nil
}
