package _463

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_islandPerimeter(t *testing.T) {
	assert.Equal(
		t,
		16,
		islandPerimeter([][]int{
			{0, 1, 0, 0},
			{1, 1, 1, 0},
			{0, 1, 0, 0},
			{1, 1, 0, 0},
		}))
	assert.Equal(
		t,
		16,
		islandPerimeter([][]int{
			{0, 1, 0, 0},
			{1, 0, 1, 0},
			{0, 1, 0, 0},
			{0, 0, 0, 0},
		}))
}
