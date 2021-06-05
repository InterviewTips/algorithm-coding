package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindRepeatNumber(t *testing.T) {
	assert.Contains(t, []int{2, 3}, findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
}
