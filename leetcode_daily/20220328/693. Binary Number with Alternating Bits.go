package _0220328

func hasAlternatingBits(n int) bool {
	var flag = n & 1
	for n != 0 {
		value := n >> 1 & 1  // 取出右移的一位
		if flag^value != 1 { // 异或运算 01 10
			return false
		}
		flag = value // 赋值为右移的一位
		n = n >> 1
	}

	return true
}
