package _207

import (
	"fmt"
)

// 维护 map
func uniqueOccurrences(arr []int) bool {
	res := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		_, ok := res[arr[i]]
		if !ok {
			res[arr[i]] = 1
		} else {
			res[arr[i]] += 1
		}
	}
	fmt.Printf("map is %v\n", res)
	res1 := make(map[int]int)
	for k, v := range res {
		_, ok := res1[v]
		if ok { // 说明存在 出现次数一致的数字
			return false
		} else {
			res1[v] = k
		}
	}
	return true
}
