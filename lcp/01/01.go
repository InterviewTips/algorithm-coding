package main

func game(guess []int, answer []int) int {
	ac := 0
	for i := 0; i < len(guess); i++ {
		if guess[i] == answer[i] {
			ac += 1
		}
	}
	return ac
}
