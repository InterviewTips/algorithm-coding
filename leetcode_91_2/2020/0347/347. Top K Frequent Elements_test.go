package _347

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	assert.Equal(t, []int{1, 2}, topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
