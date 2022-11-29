package _0221129

func minOperations(s string) int {
	// 计算最小操作的次数 要么是 010 要么是 101

	// 开头为 '0' 的情况
	count0 := 0
	for i := 0; i < len(s); i++ {
		item := s[i]
		if i&1 == 0 && item == '1' {
			count0++
		}
		if i&1 == 1 && item == '0' {
			count0++
		}
	}

	// 开头为 '1' 的情况
	count1 := 0
	for i := 0; i < len(s); i++ {
		item := s[i]
		if i&1 == 0 && item == '0' {
			count1++
		}
		if i&1 == 1 && item == '1' {
			count1++
		}
	}

	if count0 < count1 {
		return count0
	}

	return count1
}
