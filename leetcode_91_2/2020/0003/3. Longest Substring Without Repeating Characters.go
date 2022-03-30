package _003

func lengthOfLongestSubstring(s string) int {
	//init
	duplicate := make(map[byte]int)
	start, end := 0, 0
	res := 0
	for i := 0; i < len(s); i++ {
		v, ok := duplicate[s[i]]
		if ok && v >= start { // 有效重复元素
			//add to res
			res = maxValue(res, end-start+1)
			//set start
			start = v + 1
		}
		duplicate[s[i]] = i
		end = i
	}
	if len(s)-start > res {
		res = len(s) - start
	}
	return res
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring1(s string) int {
	//init
	duplicate := make(map[byte]int)
	start, end := 0, 0
	res := 0
	for i := 0; i < len(s); i++ {
		v, ok := duplicate[s[i]]
		if ok { // 有效重复元素
			//add to res
			res = maxValue(res, end-start+1)
			//clean elements
			for j := start; j < v+1; j++ {
				delete(duplicate, s[j])
			}
			//set start
			start = v + 1
		}
		duplicate[s[i]] = i
		end = i
	}
	if len(duplicate) > res {
		res = len(duplicate)
	}
	return res
}
