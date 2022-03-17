package _0220317

import "sort"

func longestWord(words []string) string {
	sort.SliceStable(words, func(i, j int) bool {
		// words[i] > words[j] 字典序降序 这样后面可以覆盖
		return len(words[i]) < len(words[j]) || (len(words[i]) == len(words[j]) && words[i] > words[j])
	})

	candidates := map[string]struct{}{
		"": {},
	}

	var longest string
	for i := 0; i < len(words); i++ {
		item := words[i]
		if _, ok := candidates[item[:len(item)-1]]; ok {
			longest = item
			candidates[item] = struct{}{}
		}
	}

	return longest
}
