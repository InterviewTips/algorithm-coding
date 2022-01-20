package two_pointers

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i] ^= s[j] // a ^ b
		s[j] ^= s[i] // b ^ a ^ b
		s[i] ^= s[j] //  a ^ a ^ b
	}

	return
}
