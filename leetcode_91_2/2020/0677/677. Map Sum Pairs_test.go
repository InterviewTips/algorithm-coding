package _677

import "testing"

func TestMapSum(t *testing.T) {
	mapSum := Constructor()
	mapSum.Insert("apple", 3)
	println(mapSum.Sum("ap")) // return 3 (apple = 3)
	mapSum.Insert("app", 2)
	println(mapSum.Sum("ap")) // return 5 (apple + app = 3 + 2 = 5)
}
