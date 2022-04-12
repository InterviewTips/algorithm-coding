package _0220412

func numberOfLines(widths []int, s string) []int {
	line := 1
	count := 0
	for i := 0; i < len(s); i++ {
		width := widths[s[i]-'a']
		if count+width > 100 {
			line++
			count = width
		} else {
			count += width
		}
	}

	return []int{line, count}

}
