package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestMissNumber(t *testing.T) {
	t.Run("首", func(t *testing.T) {
		m := missingNumber([]int{1, 2, 3})
		assert.Equal(t, 0, m, 0)
	})
	t.Run("尾", func(t *testing.T) {
		m := missingNumber([]int{0, 1, 2})
		assert.Equal(t, 3, m, 3)
	})
	t.Run("中间", func(t *testing.T) {
		m := missingNumber([]int{0, 1, 3})
		assert.Equal(t, 2, m, 2)
	})
}

func TestMissNumber1(t *testing.T) {
	t.Run("首", func(t *testing.T) {
		m := missingNumber1([]int{1, 2, 3})
		assert.Equal(t, 0, m, 0)
	})
	t.Run("尾", func(t *testing.T) {
		m := missingNumber1([]int{0, 1, 2})
		assert.Equal(t, 3, m, 3)
	})
	t.Run("中间", func(t *testing.T) {
		m := missingNumber1([]int{0, 1, 3})
		assert.Equal(t, 2, m, 2)
	})
}

func TestMissNumber2(t *testing.T) {
	t.Run("首", func(t *testing.T) {
		m := missingNumber2([]int{1, 2, 3})
		assert.Equal(t, 0, m, 0)
	})
	t.Run("尾", func(t *testing.T) {
		m := missingNumber2([]int{0, 1, 2})
		assert.Equal(t, 3, m, 3)
	})
	t.Run("中间", func(t *testing.T) {
		m := missingNumber2([]int{0, 1, 3})
		assert.Equal(t, 2, m, 2)
	})
}
