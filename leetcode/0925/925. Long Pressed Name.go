package _925

func isLongPressedName(name string, typed string) bool {
	if len(typed) < len(name) {
		return false
	}
	n := 0 // name
	for i := 0; i < len(typed); i++ {
		if typed[i] == name[n] {
			n++ // name increase 1
			if n == len(name) {
				for j := i + 1; j < len(typed); j++ {
					if typed[j] != name[n-1] {
						return false
					}
				}
				return true
			}
		} else {
			if i == 0 {
				return false
			}
			if typed[i] != typed[i-1] {
				return false
			}
		}
	}
	return false
}

// 双指针
func isLongPressedNamePlanB(name string, typed string) bool {
	i, j := 0, 0
	for j < len(typed) {
		if i < len(name) && name[i] == typed[j] {
			i++
			j++
		} else if j > 0 && typed[j] == typed[j-1] {
			j++
		} else {
			return false
		}
	}
	return i == len(name)
}
