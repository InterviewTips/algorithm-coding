package _008

import (
	"bytes"
	"math"
	"strconv"
)

func myAtoi(numStr string) int {
	flag := false
	poss := false
	neg := false
	numFlag := false
	var num bytes.Buffer
	for i := 0; i < len(numStr); i++ {
		if numStr[i] == ' ' && !flag {
			continue
		}
		if numStr[i] == '-' && !poss && !neg && !numFlag {
			poss = true
			flag = true
			continue
		}
		if numStr[i] == '+' && !flag && !poss && !numFlag {
			neg = true
			flag = true
			continue
		}
		if numStr[i] >= '0' && numStr[i] <= '9' {
			flag = true
			numFlag = true
			num.WriteByte(numStr[i])
		} else {
			break
		}
	}
	var n int
	if num.Len() > 0 {
		n, _ = strconv.Atoi(num.String())
	}
	if poss {
		n = -n
	}
	if n < math.MinInt32 {
		n = math.MinInt32
	}
	if n > math.MaxInt32 {
		n = math.MaxInt32
	}
	return n
}
