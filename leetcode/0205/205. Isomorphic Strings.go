package _205

func isIsomorphic(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	data := make(map[byte]byte)
	for i := 0; i < len(str1); i++ {
		value, ok := data[str1[i]]
		if !ok {
			data[str1[i]] = str2[i]
			continue
		}
		// 存在 则比较
		// 相同字符只能映射到同一个字符上
		if value != str2[i] {
			return false
		}
	}

	// 比较 data 中的 value 是否存在重复
	// 不同字符不能映射到同一个字符上
	data1 := make(map[byte]struct{})
	for _, value := range data {
		_, ok := data1[value]
		if ok {
			return false
		}
		data1[value] = struct{}{}
	}

	return true
}
