package leetcode

import (
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		t.Logf("%v\n", generateParenthesis(3))
	})
}
