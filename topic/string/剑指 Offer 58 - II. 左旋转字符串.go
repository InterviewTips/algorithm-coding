package string

import (
	"bytes"
)

func reverseLeftWords(source string, num int) string {
	if num > len(source) { // error
		return ""
	}

	var buf bytes.Buffer
	needReverse := source[:num]
	buf.WriteString(source[num:])
	buf.WriteString(needReverse)

	return buf.String()
}
