package main

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestExist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"

	assert.Equal(t, true, exist(board, word))

}
