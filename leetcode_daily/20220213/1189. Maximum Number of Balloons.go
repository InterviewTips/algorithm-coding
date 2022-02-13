package _0220213

func maxNumberOfBalloons(text string) int {
	// "balloon"
	data := make(map[byte]int)
	for i := 0; i < len(text); i++ {
		switch text[i] {
		case 'b', 'a', 'l', 'o', 'n':
			data[text[i]]++
		}
	}
	// 判断
	count := 0
	for {
		data['b']--
		data['a']--
		data['l'] -= 2
		data['o'] -= 2
		data['n']--
		if data['b'] >= 0 &&
			data['a'] >= 0 &&
			data['l'] >= 0 &&
			data['o'] >= 0 &&
			data['n'] >= 0 {
			count++
		} else {
			break
		}
	}

	return count
}
