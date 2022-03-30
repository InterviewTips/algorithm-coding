package _438

func findAnagrams(s string, p string) []int {
	// 滑动窗口
	// 计算长度
	pLen := len(p)
	if len(s) < pLen {
		return nil
	}

	pMapOrigin := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		pMapOrigin[p[i]] += 1
	}
	pMap := make(map[byte]int)
	// init
	for i := 0; i < pLen; i++ {
		pMap[s[i]] += 1
	}
	res := make([]int, 0)
	for i := 0; i < len(s)-pLen+1; i++ {
		if judge(pMap, pMapOrigin) {
			res = append(res, i)
		}
		// 不能直接 delete 这个 key，因为有可能出现了两次 直接删除就没了 其实也可以优化一下 为 0 则删除
		pMap[s[i]] -= 1
		if i+pLen < len(s) {
			pMap[s[i+pLen]] += 1
		}
	}

	return res
}

//判断是否一致
func judge(pMap map[byte]int, pMapOrigin map[byte]int) bool {

	for k := range pMapOrigin {
		if pMapOrigin[k] != pMap[k] {
			return false
		}
	}
	return true
}
