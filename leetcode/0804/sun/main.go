package main

import (
	"fmt"
	"strings"
)

func getCode(words []string) string {
	codeMap := MorseCodeParse()
	res := make([]string, 0)
	for i := 0; i < len(words); i++ {
		if words[i] == "" {
			res = append(res, " ")
			continue
		}
		res = append(res, string(codeMap[words[i]]))
	}

	return strings.Join(res, "")
}

func MorseCodeParse() map[string]rune {

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

	//fmt.Println(lencode))
	// 建立映射
	codeMap := make(map[string]rune)

	start := 'a'
	//fmt.Println(string(start + 25))
	for i := 0; i < len(code); i++ {
		codeMap[code[i]] = start + int32(i)
	}

	return codeMap

}

func main() {
	fmt.Println(getCode([]string{
		"...",
		"..-",
		".--.",
		".",
		".-.",
		"", // 空格
		"...",
		"---",
		".-..",
		"...",
		"-",
		"..",
		"-.-.",
		".",
	}))
	// 克卜勒专辑摩斯电码解密
	//output: super solstice
}
