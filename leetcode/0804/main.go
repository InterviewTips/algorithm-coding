package main

import (
	"fmt"
	"strings"
)

func uniqueMorseRepresentations(words []string) int {
	if len(words) == 0 {
		return 0
	}
	codeMap := MorseCodeParse()
	res := make([]string, 0)
	var solution func(word string) string
	solution = func(word string) string {
		res := make([]string, 0)
		for i := 0; i < len(word); i++ {
			//fmt.Println(string(rune(word[i])))
			res = append(res, codeMap[rune(word[i])])
		}

		return strings.Join(res, "")
	}
	for i := 0; i < len(words); i++ {
		res = append(res, solution(words[i]))
	}

	//fmt.Println(res)
	// 去除重复的元素
	lastRes := make([]string, 0)
	lastResMap := make(map[string]int)
	for i := 0; i < len(res); i++ {
		if lastResMap[res[i]] == 1 {
			continue
		}
		lastResMap[res[i]] = 1
		lastRes = append(lastRes, res[i])
	}
	return len(lastRes)
}

func MorseCodeParse() map[rune]string {

	code := []string{
		".-",
		"-...",
		"-.-.",
		"-..",
		".",
		"..-.",
		"--.",
		"....",
		"..",
		".---",
		"-.-",
		".-..",
		"--",
		"-.",
		"---",
		".--.",
		"--.-",
		".-.",
		"...",
		"-",
		"..-",
		"...-",
		".--",
		"-..-",
		"-.--",
		"--..",
	}

	fmt.Println(len(code))
	// 建立映射
	codeMap := make(map[rune]string)

	start := 'a'
	//fmt.Println(string(start + 25))
	for i := 0; i < len(code); i++ {
		codeMap[start+int32(i)] = code[i]
	}

	return codeMap

}

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}
