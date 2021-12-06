package string

import (
	"bytes"
)

// 如果剩余字符少于 k 个，则将剩余字符全部反转。
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
func reverseStr(s string, k int) string {
	var buf bytes.Buffer
	lenS := len(s)
	num := lenS / (2 * k)
	for i := 0; i < num; i++ {
		kFirst := []byte(s[2*k*i : 2*k*i+k])    // 前 k 个
		kLast := []byte(s[2*k*i+k : 2*k*(i+1)]) // 后 k 个
		buf.Write(reverseSingleStr(kFirst))
		buf.Write(kLast)
	}

	// 假定还有剩余的
	mod := lenS % (2 * k)
	if mod < k {
		buf.Write(reverseSingleStr([]byte(s[2*k*num:])))
	} else {
		buf.Write(reverseSingleStr([]byte(s[2*k*num : 2*k*num+k]))) // 前 k 个
		buf.Write([]byte(s[2*k*num+k:]))
	}

	return buf.String()
}

func reverseSingleStr(s []byte) []byte {
	// fmt.Println(s)
	for i := 0; i < len(s)/2; i++ {
		tmp := s[i]
		s[i] = s[len(s)-(i+1)]
		s[len(s)-(i+1)] = tmp
	}

	//fmt.Println(s)
	return s
}
