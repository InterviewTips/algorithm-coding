package _064

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestMinPathSum(t *testing.T) {
	assert.Equal(t, 7, minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}), "must be 7")
	assert.Equal(t, 6, minPathSum([][]int{
		{1, 2, 5},
		{3, 2, 1},
	}), "must be 6")
}
