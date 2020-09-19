package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestGame(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		assert.Equal(t,
			3,
			game([]int{1, 2, 3}, []int{1, 2, 3}),
			"must be 3",
		)
	})
	t.Run("only 1", func(t *testing.T) {
		assert.Equal(t,
			1,
			game([]int{3, 2, 2}, []int{1, 2, 3}),
			"must be 1",
		)
	})
}
