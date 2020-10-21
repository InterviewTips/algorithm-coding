package _435

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_mostIntervals(t *testing.T) {
	// [[1,2],[2,3],[3,4],[1,3]]
	var intervals [][]int
	intervals = append(intervals, []int{1, 2})
	intervals = append(intervals, []int{2, 3})
	intervals = append(intervals, []int{3, 4})
	intervals = append(intervals, []int{1, 3})
	mostIntervals(intervals)
}

func Test_eraseOverlapIntervals(t *testing.T) {
	assert.Equal(t, 1, eraseOverlapIntervals([][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 3},
	}))
	assert.Equal(t, 2, eraseOverlapIntervals([][]int{
		{1, 2},
		{1, 2},
		{1, 2},
	}))
	assert.Equal(t, 0, eraseOverlapIntervals([][]int{
		{1, 2},
		{2, 3},
	}))
}
