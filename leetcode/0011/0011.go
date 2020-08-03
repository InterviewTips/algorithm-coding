package main

import "fmt"

func maxArea(height []int) int {
	var i, j int
	mArea := getArea(height[0], height[1], 0, 1)
	for i = 0; i < len(height)-1; i++ {
		for j = i + 1; j < len(height); j++ {
			if getArea(height[i], height[j], i, j) > mArea {
				mArea = getArea(height[i], height[j], i, j)
			}
		}
	}

	return mArea

}

// 求面积
func getArea(x, y, i, j int) int {
	return mathAbs(i, j) * mathMin(x, y)
}

// 获取绝对值
func mathAbs(x, y int) int {
	if (x - y) > 0 {
		return x - y
	}
	return y - x
}

// 获取最小值
func mathMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
