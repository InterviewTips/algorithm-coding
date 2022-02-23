package _0220223

func reverseOnlyLetters(s string) string {
	left := 0
	right := len(s) - 1
	res := make([]byte, len(s))
	for left <= right {
		if noLetter(s[left]) {
			res[left] = s[left]
			left++
			continue
		}
		if noLetter(s[right]) {
			res[right] = s[right]
			right--
			continue
		}
		res[left], res[right] = s[right], s[left]
		left++
		right--
	}

	return string(res)
}

func noLetter(b byte) bool {
	if 'a' <= b && b <= 'z' ||
		'A' <= b && b <= 'Z' {
		return false
	}

	return true
}
