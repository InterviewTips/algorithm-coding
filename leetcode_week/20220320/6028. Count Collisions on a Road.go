package _0220320

func countCollisions(data string) int {
	// RL RS SL
	count := 0
	directions := []byte(data)
	// 从头到尾
	for i := 0; i < len(directions); i++ {
		item := directions[i]
		switch item {
		case 'L':
			if i > 0 && directions[i-1] == 'R' {
				count += 2
				directions[i] = 'S'
				directions[i-1] = 'S'
				continue
			}
			if i > 0 && directions[i-1] == 'S' {
				count += 1
				directions[i] = 'S'
			}
		case 'S':
			if i > 0 && directions[i-1] == 'R' {
				count += 1
				directions[i-1] = 'S'
			}
		}
	}

	// 从尾到头
	for i := len(directions) - 2; i >= 0; i-- {
		item := directions[i]
		switch item {
		case 'R':
			if directions[i+1] == 'L' {
				count += 2
				directions[i] = 'S'
				directions[i+1] = 'S'
				continue
			}
			if directions[i+1] == 'S' {
				count += 1
				directions[i] = 'S'
			}
		case 'S':
			if directions[i+1] == 'L' {
				count += 1
				directions[i+1] = 'S'
			}
		}
	}

	return count
}
