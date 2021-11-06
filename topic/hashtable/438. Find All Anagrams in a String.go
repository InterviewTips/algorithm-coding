package hashtable

import "log"

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return make([]int, 0)
	}
	// 滑动窗口 + hash table
	sMap := make(map[uint8]int)
	pMap := make(map[uint8]int)
	for i := 0; i < len(p); i++ {
		pMap[p[i]]++
	}

	res := make([]int, 0)
	// 需要减掉对应字符串长度 固定长度滑动
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
		index := i - len(p)
		if index >= 0 { // 如果满足 说明长度已经超出，这个时候需要移除之前的元素 --
			sMap[s[index]]--
			if sMap[s[index]] == 0 { // 理论上可要可不要
				delete(sMap, s[index])
			}
		}
		if isWord(sMap, pMap) {
			log.Println(sMap, pMap)
			res = append(res, i-len(p)+1)
		}
	}

	return res
}

func isWord(sMap, pMap map[uint8]int) bool {
	for k, v := range pMap {
		if sMap[k] < v {
			return false
		}
	}

	return true
}
