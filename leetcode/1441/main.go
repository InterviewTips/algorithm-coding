package main

func buildArray(target []int, n int) []string {
	res := make([]string, 0)
	if n < len(target) {
		return res
	}
	start := 0
	for i := 1; i <= n; i++ {
		if target[start] == i {
			res = append(res, "Push")
			start++
			if start >= len(target) {
				return res
			}
		} else {
			res = append(res, "Push")
			res = append(res, "Pop")
		}
	}

	return res
}
