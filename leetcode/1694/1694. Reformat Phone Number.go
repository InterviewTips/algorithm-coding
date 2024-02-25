package _694

import (
	"bytes"
	"strings"
)

func reformatNumber(number string) string {
	number = strings.Replace(number, " ", "", -1)
	number = strings.Replace(number, "-", "", -1)

	var res bytes.Buffer

	if len(number)%3 == 1 { // 最后是 xx-xx
		for i := 0; i < len(number)-4; i++ {
			if i%3 == 0 && i != 0 {
				res.WriteByte('-')
			}
			res.WriteByte(number[i])
		}
		if len(res.String()) > 0 {
			res.WriteByte('-')
		}
		res.WriteString(number[len(number)-4 : len(number)-2])
		res.WriteByte('-')
		res.WriteString(number[len(number)-2:])
		return res.String()
	}
	if len(number)%3 == 2 { // 最后是 xx
		for i := 0; i < len(number)-2; i++ {
			if i%3 == 0 && i != 0 {
				res.WriteByte('-')
			}
			res.WriteByte(number[i])
		}
		if len(res.String()) > 0 {
			res.WriteByte('-')
		}
		res.WriteString(number[len(number)-2:])
		return res.String()
	}

	for i := 0; i < len(number); i++ {
		if i%3 == 0 && i != 0 {
			res.WriteByte('-')
		}
		res.WriteByte(number[i])
	}

	return res.String()
}
