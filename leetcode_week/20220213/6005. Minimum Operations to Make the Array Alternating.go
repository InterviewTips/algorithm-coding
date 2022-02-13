package _0220213

import (
	"sort"
)

func minimumOperations(nums []int) int {
	if len(nums) <= 1 { // 本来就是
		return 0
	}
	// 统计 map
	num1Data := make(map[int]int)
	num2Data := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		num1Data[nums[i]]++
		i++
		if i < len(nums) {
			num2Data[nums[i]]++
		}
	}
	num1s := make([][2]int, 0)
	for k, v := range num1Data {
		num1s = append(num1s, [2]int{k, v})
	}
	num2s := make([][2]int, 0)
	for k, v := range num2Data {
		num2s = append(num2s, [2]int{k, v})
	}
	// 按照顺序排好
	sort.SliceStable(num1s, func(i, j int) bool {
		return num1s[i][1] > num1s[j][1] // 这里的 i, j 容易混淆 一不小心可能写错
	})
	sort.SliceStable(num2s, func(i, j int) bool {
		return num2s[i][1] > num2s[j][1]
	})
	// 取出 num1 num2
	var num1 [2]int
	var num2 [2]int

	num1 = num1s[0]
	num2 = num2s[0]
	if num1[0] != num2[0] { // 不同 直接返回
		return len(nums) - (num1[1] + num2[1])
	}

	// 取出次之的数
	var x int
	var y int
	index := 1
	if index < len(num1s) {
		x = num2[1] + num1s[index][1]
	} else {
		x = num2[1] + 0 // 次之的数 出现的次数为 0
	}
	if index < len(num2s) {
		y = num1[1] + num2s[index][1]
	} else {
		y = num1[1] + 0 // 次之的数 出现的次数为 0
	}

	count := len(nums) - max(x, y)

	return count
}

func max(x int, y int) int {
	if x > y {
		return x
	}

	return y
}
