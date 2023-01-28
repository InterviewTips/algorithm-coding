package _0230128

func waysToMakeFair(nums []int) int {
	// https://leetcode.cn/problems/ways-to-make-a-fair-array/solution/by-lcbin-3jl3/
	// t1 奇数和
	// t2 偶数和
	// s1 全部奇数和
	// s2 全部偶数和
	var t1, t2, s1, s2 int
	for i := 0; i < len(nums); i++ {
		if i&1 == 1 {
			s1 += nums[i]
		} else {
			s2 += nums[i]
		}
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		if i&1 == 0 && t1+s2-t2-nums[i] == t2+s1-t1 {
			count++
		}
		if i&1 == 1 && t2+s1-t1-nums[i] == t1+s2-t2 {
			count++
		}
		if i&1 == 1 {
			t1 += nums[i]
		} else {
			t2 += nums[i]
		}
	}

	return count
}

// timeout
func calcTimeout(left []int, right []int) bool {
	//log.Println("left", left, "right", right)
	sum1 := 0
	sum2 := 0

	if len(left) > 0 {
		for i := 0; i < len(left); i++ {
			if i&1 == 1 { // 奇数
				sum2 += left[i]
			} else {
				sum1 += left[i]
			}
		}
	}
	if len(right) > 0 && len(left)&1 == 0 {
		for i := 0; i < len(right); i++ {
			if i&1 == 1 { // 奇数
				sum2 += right[i]
			} else {
				sum1 += right[i]
			}
		}
	}
	if len(right) > 0 && len(left)&1 == 1 {
		for i := 0; i < len(right); i++ {
			if i&1 == 0 {
				sum2 += right[i]
			} else {
				sum1 += right[i]
			}
		}
	}

	//log.Println("sum1", sum1, "sum2", sum2)

	return sum1 == sum2
}
