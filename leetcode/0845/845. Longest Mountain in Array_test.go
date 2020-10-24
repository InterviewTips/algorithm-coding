package _845

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_longestMountain(t *testing.T) {
	assert.Equal(t, 5, longestMountain([]int{2, 1, 4, 7, 3, 2, 5}))
	assert.Equal(t, 0, longestMountain([]int{2, 2, 2}))
}
