package backtracking

import "log"

func partition(s string) [][]string {
	res := make([][]string, 0)
	path := make([]string, 0)
	var backtracking func(sVar string)
	backtracking = func(sVar string) {
		//log.Println( "path is", path)
		if getPathLen(path) == len(s) {
			value := make([]string, len(path))
			copy(value, path)
			res = append(res, value)
			return
		}

		for i := 0; i < len(sVar); i++ {
			newValue := sVar[:i+1]
			if !isPalindrome(newValue) { // 不是回文字符串
				continue
			}
			newSVar := sVar[i+1:] // 不断切割 传入新的 str value
			path = append(path, newValue)
			backtracking(newSVar)
			path = path[:len(path)-1]
		}
	}

	backtracking(s)

	log.Println(res)
	return res
}

func isPalindrome(data string) bool {
	for i := 0; i < len(data)/2; i++ {
		if data[i] != data[len(data)-1-i] {
			return false
		}
	}

	return true
}

func getPathLen(path []string) int {
	sum := 0
	for i := 0; i < len(path); i++ {
		sum += len(path[i])
	}

	return sum
}

// var partition = partitionStandard
func partitionStandard(s string) [][]string {
	res := make([][]string, 0)
	path := make([]string, 0)
	var backtracking func(index int)
	backtracking = func(index int) {
		//log.Println("index is", index, "path is", path)
		// index 表示开始切割的位置，如果 >= len(s) 说明已经来到了末尾
		if index == len(s) {
			value := make([]string, len(path))
			copy(value, path)
			res = append(res, value)
			return
		}

		for i := index; i < len(s); i++ {
			newValue := s[index : i+1]
			if !isPalindrome(newValue) { // 不是回文字符串
				continue
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
