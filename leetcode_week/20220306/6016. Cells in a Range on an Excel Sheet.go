package _0220306

import "fmt"

func cellsInRange(s string) []string {
	col1 := s[0]
	row1 := s[1]
	col2 := s[3]
	row2 := s[4]

	res := make([]string, 0)
	for c := col1; c <= col2; c++ {
		for r := row1; r <= row2; r++ {
			res = append(res, fmt.Sprintf("%c%c", c, r))
		}
	}

	return res
}
