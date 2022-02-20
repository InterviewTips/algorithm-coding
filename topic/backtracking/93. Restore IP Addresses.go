package backtracking

import (
	"log"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	if len(s) > 12 { // 剪枝
		return res
	}
	path := make([]string, 0)
	var backtracking func(index int)
	backtracking = func(index int) {
		if len(path) == 3 {
			// 取出 index 之后的
			if index >= len(s) {
				return
			}
			value := s[index:]
			if !isValidIP(value) {
				return
			}
			path = append(path, value)
			res = append(res, strings.Join(path, "."))
			path = path[:len(path)-1] // 回撤
			return
		}

		for i := index; i < len(s); i++ {
			newValue := s[index : i+1]
			if !isValidIP(newValue) {
				// continue
				// 其实后面横向的都不合法了 可以直接 break
				break
			}
			path = append(path, newValue)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtracking(0)

	log.Println(res)
	return res
}

func isValidIP(data string) bool {
	if len(data) == 0 {
		return false
	}
	if data == "0" {
		return true
	}
	if len(data) > 1 && data[0] == '0' {
		return false
	}
	value, err := strconv.Atoi(data)
	if err != nil {
		log.Println(err)
		return false
	}
	if 0 < value && value <= 255 {
		return true
	}

	return false
}
