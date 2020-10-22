package _763

func partitionLabels(S string) []int {
	if len(S) == 0 {
		return nil
	}
	//计算所有字符的最后一次出现的位置
	m := make(map[byte]int)
	for i := 0; i < len(S); i++ {
		m[S[i]] = i
	}
	//双指针
	start := 0
	end := 0
	var res []int
	for i := 0; i < len(S); i++ {
		if m[S[i]] > end {
			end = m[S[i]]
		}
		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}
	return res
}
