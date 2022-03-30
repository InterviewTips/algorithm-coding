package _0220329

func maxConsecutiveChar(answerKey string, k int, ch byte) (ans int) {
	left, sum := 0, 0
	for right := 0; right < len(answerKey); right++ {
		if answerKey[right] != ch {
			sum++
		}
		for sum > k {
			if answerKey[left] != ch { // 说明是转换而来
				sum--
			}
			left++ // left 右移动
		}
		ans = max(ans, right-left+1)
	}
	return
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	// 比较 T 和 F 最大的是哪个
	return max(maxConsecutiveChar(answerKey, k, 'T'),
		maxConsecutiveChar(answerKey, k, 'F'))
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
