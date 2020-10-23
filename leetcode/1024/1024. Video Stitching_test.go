package _024

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_clips_sort(t *testing.T) {
	videoStitching([][]int{
		{0, 2},
		{0, 8},
		{4, 6},
		{8, 10},
		{1, 7},
		{1, 9},
		{1, 5},
		{5, 9},
	}, 0)
}

func Test_videoStitching(t *testing.T) {
	assert.Equal(
		t,
		3,
		videoStitching([][]int{
			{0, 2},
			{4, 6},
			{8, 10},
			{1, 9},
			{1, 5},
			{5, 9},
		}, 10),
	)
	assert.Equal(
		t,
		-1,
		videoStitching([][]int{
			{0, 1},
			{1, 2},
		}, 5),
	)
	assert.Equal(
		t,
		2,
		videoStitching([][]int{
			{0, 4}, {2, 8},
		}, 5),
	)
}
