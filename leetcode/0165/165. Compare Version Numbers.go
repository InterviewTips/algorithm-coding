package _165

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	version1s := strings.Split(version1, ".")
	version2s := strings.Split(version2, ".")

	for i, j := 0, 0; i < len(version1s) && j < len(version2s); i, j = i+1, j+1 {
		num1, _ := strconv.Atoi(version1s[i])
		num2, _ := strconv.Atoi(version2s[j])
		if num1 > num2 {
			return 1
		}
		if num1 < num2 {
			return -1
		}
		// 等于则继续比较
	}

	if len(version1s) == len(version2s) {
		return 0
	}

	// 可能有长度不一样的时候
	if len(version1s) > len(version2s) {
		for i := len(version2s); i < len(version1s); i++ {
			num, _ := strconv.Atoi(version1s[i])
			if num > 0 { // version1 > version2
				return 1
			}
		}
	}

	if len(version2s) > len(version1s) {
		for i := len(version1s); i < len(version2s); i++ {
			num, _ := strconv.Atoi(version2s[i])
			if num > 0 { // version2 > version1
				return -1
			}
		}
	}

	return 0
}
