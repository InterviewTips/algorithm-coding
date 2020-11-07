package _769

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_maxChunksToSorted(t *testing.T) {
	assert.Equal(
		t,
		1,
		maxChunksToSorted([]int{4, 3, 2, 1, 0}),
	)
	assert.Equal(
		t,
		4,
		maxChunksToSorted([]int{1, 0, 2, 3, 4}),
	)
}
