package main

// 找到旋转数组的最小数字
func minArray(numbers []int) int {
	if len(numbers) == 0 {
		panic("error")
	}

	res := numbers[0]
	// 递增排序，找出第一个逆序的数字
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < res {
			return numbers[i]
		} else {
			res = numbers[i]
		}
	}

	// 没有找到，所以都是递增
	return numbers[0]
}
