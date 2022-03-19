package leetcode_interview

func replaceSpaces(S string, length int) string {
	res := make([]byte, 0)
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			res = append(res, "%20"...)
			continue
		}

		res = append(res, S[i])
	}

	return string(res)
}
