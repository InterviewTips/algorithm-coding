package _0220517

import (
	"sort"
)

func isAlienSorted(words []string, order string) bool {
	// words 需要根据 order 进行正确排序
	// orderMap
	orderMap := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		orderMap[order[i]] = i
	}

	pre := make([]string, len(words))
	copy(pre, words)

	sort.Slice(words, func(a, b int) bool {
		// 进一步比较 words[a] and words[b]
		lenA := len(words[a])
		lenB := len(words[b])
		var length = lenA
		if lenA > lenB {
			length = lenB
		}

		for i := 0; i < length; i++ {
			valueA := orderMap[words[a][i]]
			valueB := orderMap[words[b][i]]
			if valueA == valueB {
				continue
			}
			if valueA > valueB {
				return false
			}
			if valueA < valueB {
				return true
			}
		}
		// 还没对比完 说明都相等
		if lenA < lenB { // 长度小
			return true
		}

		return false
	})

	for i := 0; i < len(pre); i++ {
		if pre[i] != words[i] {
			return false
		}
	}

	return true
}
