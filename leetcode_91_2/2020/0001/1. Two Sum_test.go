package _001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	assert.Equal(
		t,
		[]int{0, 1},
		twoSum([]int{2, 7, 11, 15}, 9),
		"must be [0, 1]",
	)
}
