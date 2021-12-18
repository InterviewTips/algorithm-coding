package string

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	a := strings.Split(s, " ")
	b := make([]string, 0)
	for i := 0; i < len(a); i++ {
		if a[i] == "" {
			continue
		}
		b = append(b, a[i])
	}
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-(i+1)] = b[len(b)-(i+1)], b[i]
	}
	fmt.Println(b)
	c := strings.Join(b, " ")
	// fmt.Println(b)
	return c
}

// 分三步
// 1、移除多余空格 快慢指针
// 2、整个字符串反转
// 3、单词反转
func reverseWords1(s string) string {
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
	// 移除中间的多余空格
	for ; fastIndex < len(s); fastIndex++ {
		if fastIndex-1 > 0 && s[fastIndex-1] == s[fastIndex] && s[fastIndex] == ' ' {
			continue
		} else {
			res[slowIndex] = s[fastIndex]
			slowIndex++
		}
	}
	// 移除 res 后面的多余空格
	// 此时 slowIndex 是数组的长度
	if slowIndex-1 > 0 && res[slowIndex-1] == ' ' { // 判断一下最后一个字符是不是空格 注意这里是 res 的 [slowIndex-1]
		return res[:slowIndex-1]
	} else {
		return res[:slowIndex]
	}
}

func reverseSingleWord(reverseBytes []byte) string {
	res := make([]string, 0)
	for i := 0; i < len(reverseBytes); i++ {
		j := i
		for j < len(reverseBytes) && reverseBytes[j] != ' ' {
			j++
		}
		res = append(res, string(reverseSingleStr(reverseBytes[i:j])))
		i = j // i 变为 j 这个时候 j 是空格的索引
	}

	return strings.Join(res, " ")
}
