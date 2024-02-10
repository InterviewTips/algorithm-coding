package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinArray(t *testing.T) {
	assert.Equal(t, 1, minArray([]int{3, 4, 5, 1, 2}))
	assert.Equal(t, 0, minArray([]int{2, 2, 2, 0, 1}))
}
