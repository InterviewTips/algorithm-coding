package _447

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfBoomerangs(t *testing.T) {
	// [0,0],[1,0],[2,0]
	assert.Equal(t, 2,
		numberOfBoomerangs([][]int{
			{0, 0},
			{1, 0},
			{2, 0},
		}),
		"must be 2",
	)
}
