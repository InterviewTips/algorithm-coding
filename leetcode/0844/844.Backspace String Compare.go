package _844

const flag = byte('#')

func backspaceCompare(S string, T string) bool {
	var (
		a []byte
		b []byte
	)
	s := []byte(S)
	t := []byte(T)
	for i := 0; i < len(s); i++ {
		if s[i] == flag { // 回退
			if len(a) >= 1 {
				a = a[:len(a)-1]
			}
		} else {
			a = append(a, s[i])
		}
	}
	//fmt.Printf("a is %s\n", string(a))
	for i := 0; i < len(t); i++ {
		if t[i] == flag { // 回退
			if len(b) >= 1 {
				b = b[:len(b)-1]
			}
		} else {
			b = append(b, t[i])
		}
	}
	//fmt.Printf("b is %s\n", string(b))

	return string(a) == string(b)
}
