package _024

import "fmt"

func videoStitching(clips [][]int, T int) int {
	// sort
	maxn := make([]int, T) //记录同一左端点，距离最远的右端点数组
	last, pre := 0, 0
	for i := 0; i < len(clips); i++ {
		l, r := clips[i][0], clips[i][1]
		if l < T && r > maxn[l] {
			maxn[l] = r
		}
	}
	fmt.Printf("%v\n", maxn)
	count := 0
	for i, v := range maxn {
		if v > last {
			last = v
		}
		if i == last {
			return -1
		}
		if i == pre { //区间用完了，开启下一个区间
			count++
			pre = last
		}
	}

	return count
}
