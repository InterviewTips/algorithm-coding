package _022season

func perfectMenu(materials []int, cookbooks [][]int, attribute [][]int, limit int) int {
	// materials bagWeight
	// cookbooks weight
	// attribute value

	for i := 0; i < len(cookbooks); i++ {
		for j := materials; sum(getSub(j, cookbooks[i])); j = getSub(j, cookbooks[i]) {
		}
	}

	return 0
}

func getSub(m, c []int) []int {
	for i := 0; i < len(m); i++ {
		m[i] = m[i] - c[i]
	}

	return m
}

func sum(num []int) bool {
	for i := 0; i < len(num); i++ {
		if num[i] < 0 {
			return false
		}
	}

	return true
}
