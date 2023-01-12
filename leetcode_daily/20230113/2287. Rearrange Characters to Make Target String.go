package _0230113

func rearrangeCharacters(s string, target string) int {
	// build map and then count the max times
	dataMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		dataMap[s[i]]++
	}

	count := 0
	for {
		for i := 0; i < len(target); i++ {
			dataMap[target[i]]--
			if dataMap[target[i]] < 0 {
				goto end
			}
		}
		count++
	}
end:

	return count
}
