package string

func repeatedSubstringPattern(s string) bool {
	if s == "" {
		panic("nil string")
	}

	next := getNext(s)
	if next[len(s)-1] != 0 && len(s)%(len(s)-next[len(s)-1]) == 0 {
		return true
	}

	return false
}
