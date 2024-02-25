package solution1

import "strings"

func reformatNumber(number string) string {
	number = strings.ReplaceAll(strings.ReplaceAll(number, "-", ""), " ", "")
	var format func(sub string) string
	format = func(sub string) string {
		if len(sub) == 2 {
			return sub
		}
		if len(sub) == 3 {
			return sub
		}
		if len(sub) == 4 {
			return sub[:2] + "-" + sub[2:]
		}
		return sub[:3] + "-" + format(sub[3:])
	}
	return format(number)
}
