package main

import "testing"

func TestRunningSum(t *testing.T) {
	t.Log(runningSum([]int{1, 2, 3, 4}))
	t.Log(runningSum([]int{1, 1, 1, 1, 1}))
	t.Log(runningSum([]int{3, 1, 2, 10, 1}))
}
