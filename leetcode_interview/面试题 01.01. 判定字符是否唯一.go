package leetcode_interview

func isUnique(astr string) bool {
	data := make(map[byte]struct{})
	for i := 0; i < len(astr); i++ {
		_, ok := data[astr[i]]
		if ok {
			return false
		}
		data[astr[i]] = struct{}{}
	}

	return true
}

func isUniqueAnother(astr string) bool {
	var left int64 = 0  // 高位
	var right int64 = 0 // 低位
	for i := 0; i < len(astr); i++ {
		c := astr[i]
		if c >= 64 { // int64 只有 64 位
			var bitIndex int64 = 1 << (c - 64)
			if left&bitIndex != 0 {
				return false
			}
			left |= bitIndex
		} else { // c 0~63
			var bitIndex int64 = 1 << c
			if right&bitIndex != 0 {
				return false
			}
			right |= bitIndex
		}
	}

	return true
}
