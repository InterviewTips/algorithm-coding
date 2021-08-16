package main

import "testing"

func TestMyPow(t *testing.T) {
	t.Log(myPow(2.00000, 10))
	t.Log(myPow(2.10000, 3))
	t.Log(myPow(2.00000, -2))
}
