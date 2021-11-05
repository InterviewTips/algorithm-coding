package hashtable

func canConstruct(ransomNote string, magazine string) bool {
	// 构建 map
	magazineMap := make(map[uint8]int)
	for i := 0; i < len(magazine); i++ {
		m := magazine[i]
		magazineMap[m]++
	}

	for i := 0; i < len(ransomNote); i++ {
		r := ransomNote[i]
		magazineMap[r]--
		if magazineMap[r] < 0 {
			return false
		}
	}

	return true
}
