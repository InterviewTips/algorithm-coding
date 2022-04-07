package leetcode

func rotateString(s string, goal string) bool {
	if len(s) == 1 {
		if s == goal {
			return true
		}
		return false
	}

	source := s
	for {
		// rotate
		b := append([]byte(s[1:]), s[0])
		s = string(b)
		if s == goal {
			return true
		}
		if s == source {
			break
		}
	}

	return false
}
