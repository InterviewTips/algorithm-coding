package _0220417

import (
	"bytes"
	"math"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	bannedMap := make(map[string]struct{})
	for i := 0; i < len(banned); i++ {
		bannedMap[banned[i]] = struct{}{}
	}
	data := make(map[string]int)
	paragraph = strings.ToLower(paragraph)
	var cur bytes.Buffer
	for i := 0; i < len(paragraph); i++ {
		if 'a' <= paragraph[i] && paragraph[i] <= 'z' {
			cur.WriteByte(paragraph[i])
		} else {
			if cur.String() != "" {
				_, ok := bannedMap[cur.String()]
				if !ok {
					data[cur.String()]++
				}
			}
			cur = bytes.Buffer{}
		}
	}

	// 需要再处理 cur
	if cur.String() != "" {
		_, ok := bannedMap[cur.String()]
		if !ok {
			data[cur.String()]++
		}
	}

	maxValue := math.MinInt64
	var maxStr string
	for k, v := range data {
		if v > maxValue {
			maxValue = v
			maxStr = k
		}
	}

	return maxStr

}
