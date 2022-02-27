package _0220227

func minSteps(s string, t string) int {
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}

	count := 0
	for k, v := range sMap {
		countK := v - tMap[k]
		if countK < 0 {
			countK = -countK
		}
		count += countK
		delete(tMap, k)
	}

	for k, v := range tMap {
		countK := v - sMap[k]
		if countK < 0 {
			countK = -countK
		}
		count += countK
	}

	return count
}
