package _0220314

import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	// 找出最小索引和 暴力破解
	// 都是唯一
	// 保证有答案
	// 用 map 吧
	list1Data := make(map[string]int)
	for i := 0; i < len(list1); i++ {
		list1Data[list1[i]] = i
	}
	min := math.MaxInt64
	resData := make(map[int][]string)
	for i := 0; i < len(list2); i++ {
		item := list2[i]
		v, ok := list1Data[item]
		if !ok {
			continue
		} else {
			// 计算索引
			indexSum := v + i
			if indexSum < min {
				min = indexSum
			}
			// 赋值
			_, ok = resData[indexSum]
			if !ok {
				resData[indexSum] = []string{item}
			} else {
				resData[indexSum] = append(resData[indexSum], item)
			}
		}
	}

	return resData[min]
}
