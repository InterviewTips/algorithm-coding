package greedy

func lemonadeChange(bills []int) bool {
	data := make(map[int]int)

	for i := 0; i < len(bills); i++ {
		switch bills[i] {
		case 5:
			data[5]++
		case 10:
			// 找零 5
			if data[5] > 0 {
				data[5]--
				data[10]++
			} else {
				return false
			}
		case 20:
			// 优先找 10, 5 美元更有优势正常找零
			if data[10] > 0 && data[5] > 0 {
				data[10]--
				data[5]--
				data[20]++
				continue
			}
			if data[5] >= 3 {
				data[5] -= 3
				data[20]++
				continue
			}

			return false
		default:
			panic("err bill")
		}
	}

	return true
}
