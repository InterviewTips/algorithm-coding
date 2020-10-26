package _365

import (
	"fmt"
	"sort"
)

func smallerNumbersThanCurrent1(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	var res []int
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			// nums[i] compare with nums[j]
			if i != j && nums[j] < nums[i] {
				count += 1
			}
		}
		res = append(res, count)
	}
	return res
}

type pair struct {
	v   int
	pos int // 原始位置索引
}

func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	data := make([]pair, n)
	for i, v := range nums {
		data[i] = pair{v, i}
	}
	sort.Slice(data, func(i, j int) bool { return data[i].v < data[j].v })
	ans := make([]int, n) //结果
	prev := -1
	fmt.Printf("data is %v\n", data)
	for i, d := range data {
		if prev == -1 || d.v != data[i-1].v { //d.v == data[i-1].v 则不需要更新 prev
			prev = i
		}
		ans[d.pos] = prev
	}
	return ans
}
