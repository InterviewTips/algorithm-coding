package _0220220

func isOneBitCharacter(bits []int) bool {
	i := 0
	for i < len(bits)-1 {
		if bits[i] == 0 {
			i++
		} else {
			i += 2
		}
	}
	// len(bits) - 1 必为 0
	return i == len(bits)-1
}
