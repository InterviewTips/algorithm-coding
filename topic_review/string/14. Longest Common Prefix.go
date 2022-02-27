package string

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	index := 0
	for {
		var letter byte // byte 零值是啥
		for i := 0; i < len(strs); i++ {
			if index >= len(strs[i]) {
				goto end
			}
			if letter == byte(0) {
				letter = strs[i][index]
			}
			if strs[i][index] != letter {
				goto end
			}
		}
		index++
	}
end:

	return strs[0][:index]
}
