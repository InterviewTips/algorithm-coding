package dp

import "log"

func wordBreak(s string, wordDict []string) bool {
	res := make([]bool, len(s)+1)

	res[0] = true
	// 字符串长度为 i 的话 res[i] = true，表示可以拆分为一个或多个在字典中出现的单词。
	// 注意是长度
	// 表示的是 0 ~ i-1 索引 所以 newWord 截取的时候 需要从 j 开始而不是 j+1
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ { // 物品是什么
			newWord := s[j:i] // 开头索引为 j, 到末尾 i
			if isInWordDict(wordDict, newWord) && res[j] {
				log.Println(newWord, j)
				res[i] = true
			}
		}
		log.Println(res)
	}

	return res[len(s)]
}

func isInWordDict(wordDict []string, word string) bool {
	for i := 0; i < len(wordDict); i++ {
		if word == wordDict[i] {
			return true
		}
	}

	return false
}
