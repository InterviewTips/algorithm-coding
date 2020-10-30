package _381

import (
	"testing"
)

func Test_RandomizedCollection(t *testing.T) {
	r := Constructor()
	println(r.Insert(1))
	println(r.Insert(1))
	println(r.Insert(2))
	t.Logf("get random is %v\n", r.GetRandom())
	println(r.Remove(1))
	t.Logf("get random is %v\n", r.GetRandom())
}
