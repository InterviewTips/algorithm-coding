package _859

func sumIndicesWithKSetBits(nums []int, k int) int {
	sum := 0
	for i, num := range nums {
		if bitCount(i) == k {
			sum += num
		}
	}
	return sum
}

func bitCount(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}
