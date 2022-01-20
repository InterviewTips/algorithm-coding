package two_pointers

import "strings"

// 1、移除多余空格 快慢指针
// 2、整个字符串反转
// 3、单个单词反转
func reverseWords(s string) string {
	noSpace := removeExtraSpace(s)
	reverseBytes := reverseSingleStr(noSpace)
	res := reverseSingleWord(reverseBytes)
	return res
}

func removeExtraSpace(s string) []byte {
	slowIndex := 0
	fastIndex := 0
	res := make([]byte, len(s))
	// 移除前面的空格
	for len(s) > 0 && fastIndex < len(s) && s[fastIndex] == ' ' {
		fastIndex++
	}

	// 移除中间的空格
	for ; fastIndex < len(s); fastIndex++ {
		if fastIndex-1 > 0 && s[fastIndex-1] == s[fastIndex] && s[fastIndex] == ' ' { // 前后都是空格
			continue
		} else {
			res[slowIndex] = s[fastIndex]
			slowIndex++
		}
	}

	// 移除后面的空格
	if res[slowIndex-1] == ' ' {
		return res[:slowIndex-1]
	}

	return res[:slowIndex]
}

func reverseSingleStr(s []byte) []byte {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - (i + 1)
		s[i] ^= s[j]
		s[j] ^= s[i]
		s[i] ^= s[j]
	}
	return s
}

func reverseSingleWord(reverseBytes []byte) string {
	res := make([]string, 0)
	for i := 0; i < len(reverseBytes); i++ {
		j := i
		for j < len(reverseBytes) && reverseBytes[j] != ' ' {
			j++
		}
		// 此时 j 是空格索引
		res = append(res, string(reverseSingleStr(reverseBytes[i:j])))
		i = j
	}

	return strings.Join(res, " ")
}
