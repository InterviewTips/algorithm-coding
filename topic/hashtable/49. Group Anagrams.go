package hashtable

import (
	"bytes"
	"log"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	wordsMap := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		singleS := strs[i]
		// string 不能直接排序
		//sort.SliceStable(singleS, func(i, j int) bool {
		//	return singleS[i] < singleS[j]
		//})
		singleS = stringSort(singleS)
		_, ok := wordsMap[singleS]
		if !ok {
			wordsMap[singleS] = make([]string, 0)
		}
		wordsMap[singleS] = append(wordsMap[singleS], strs[i])
	}

	res := make([][]string, 0)
	for _, v := range wordsMap {
		res = append(res, v)
	}

	return res
}

func stringSort(s string) string {
	sortRune := make([]uint8, len(s))
	for i := 0; i < len(s); i++ {
		sortRune[i] = s[i]
	}

	sort.SliceStable(sortRune, func(i, j int) bool {
		return sortRune[i] < sortRune[j]
	})

	log.Println(sortRune)

	var b bytes.Buffer
	for i := 0; i < len(sortRune); i++ {
		_, err := b.WriteRune(rune(sortRune[i]))
		if err != nil {
			log.Println(err)
		}
	}

	return b.String()

}
