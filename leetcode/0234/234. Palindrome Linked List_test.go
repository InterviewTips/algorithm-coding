package _234

import (
	"testing"

	"github.com/InterviewTips/algorithm-coding/guns"
	"github.com/bmizerany/assert"
)

func Test_isPalindrome(t *testing.T) {
	assert.Equal(
		t,
		false,
		isPalindrome(guns.GenLinkList([]int{1, 2})),
	)
	assert.Equal(
		t,
		true,
		isPalindrome(guns.GenLinkList([]int{1, 2, 2, 1})),
	)
}
