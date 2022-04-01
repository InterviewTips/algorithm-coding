package _0220401

import "sort"

func canReorderDoubled(arr []int) bool {
	// 要考虑负数 -2 -4 递减
	// 正整数 2 4 递增
	// 要考虑负数 -2 -4 递减
	// 正整数 2 4 递增

	data := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		data[arr[i]]++
	}

	if data[0]&1 == 1 { // 判断 0 的数量是否为偶数
		return false
	}

	dataArr := make([]int, 0)
	for k, _ := range data {
		dataArr = append(dataArr, k)
	}

	sort.SliceStable(dataArr, func(i, j int) bool { // 排序方式是按照绝对值来
		return getAbs(dataArr[i]) < getAbs(dataArr[j])
	})

	for i := 0; i < len(dataArr); i++ {
		if data[2*dataArr[i]] < data[dataArr[i]] { // 如果小于则有部分 dataArr[i] 找不到
			return false
		}

		data[2*dataArr[i]] -= data[dataArr[i]]
	}

	return true
}

func getAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
