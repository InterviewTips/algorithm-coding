package _348

func zeroFilledSubarray(nums []int) (ans int64) {
	c := 0
	for _, num := range nums {
		if num == 0 {
			c++
			ans += int64(c)
		} else {
			c = 0
		}
	}
	return
}
