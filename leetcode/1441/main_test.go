package main

import "testing"

func TestBuildArray(t *testing.T) {
	res := buildArray([]int{1, 3}, 3)
	t.Log(res)
}
