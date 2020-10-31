package _066

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_plusOne(t *testing.T) {
	assert.Equal(t, []int{1, 0}, plusOne([]int{9}))
	assert.Equal(t, []int{}, plusOne([]int{}))
	assert.Equal(t, []int{4, 3, 2, 2}, plusOne([]int{4, 3, 2, 1}))
}
