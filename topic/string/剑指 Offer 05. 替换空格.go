package string

func replaceSpace(s string) string {
	res := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res = append(res, '%', '2', '0')
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}
