package _678

import "bytes"

func interpret(command string) string {
	var res bytes.Buffer
	var tmp bytes.Buffer
	for i := 0; i < len(command); i++ {
		item := command[i]
		if item == 'G' {
			res.WriteByte(item)
		}
		if item == '(' {
			tmp.WriteByte(item)
			continue
		}
		if tmp.Len() > 0 && item != ')' {
			tmp.WriteByte(item)
			continue
		}
		if item == ')' {
			if tmp.String() == "(" {
				res.WriteByte('o')
				tmp.Reset()
				continue
			}
			if tmp.String() == "(al" {
				res.WriteByte('a')
				res.WriteByte('l')
				tmp.Reset()
			}
		}
	}

	return res.String()
}
