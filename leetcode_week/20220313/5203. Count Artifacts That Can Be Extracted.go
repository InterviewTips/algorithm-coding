package _0220313

func digArtifacts(_ int, artifacts [][]int, dig [][]int) int {
	data := make(map[[2]int]bool)
	for i := 0; i < len(dig); i++ {
		item := dig[i]
		data[[2]int{item[0], item[1]}] = true
	}

	//r1-r2 c1-c2
	res := 0
	//next:
	for i := 0; i < len(artifacts); i++ {
		//log.Println("i is", i)
		item := artifacts[i]
		for r := item[0]; r <= item[2]; r++ {
			for c := item[1]; c <= item[3]; c++ {
				if !data[[2]int{r, c}] {
					goto gotoNext // 两种写法
					//continue next
				}
			}
		}
		res++
	gotoNext:
	}

	return res
}
