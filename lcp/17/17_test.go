package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestCalc(t *testing.T) {
	t.Run("AB", func(t *testing.T) {
		assert.Equal(t, 4, calculate("AB"), "must be 4")
	})

	t.Run("A", func(t *testing.T) {
		assert.Equal(t, 2, calculate("A"), "must be 2")
	})

	t.Run("B", func(t *testing.T) {
		assert.Equal(t, 2, calculate("B"), "must be 2")
	})
}
