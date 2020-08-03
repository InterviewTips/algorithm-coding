package main

import "fmt"

func maxArea(height []int) int {
	var (
		i     = 0
		j     = len(height) - 1
		mArea = 0
	)
	for i < j {
		if height[i] > height[j] {
			mArea = mathMax(mArea, height[j]*(j-i))
			j--
		} else {
			mArea = mathMax(mArea, height[i]*(j-i))
			i++
		}
	}

	return mArea

}

// 获取最大值
func mathMax(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
