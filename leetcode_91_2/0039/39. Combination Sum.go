package _039

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	var comb []int
	var dfs func([]int, int, int)
	dfs = func(candidates []int, target, index int) {
		fmt.Printf("target is %v, index is %v\n", target, index)
		if target == 0 {
			fmt.Printf("符合条件 comb, %v\n", comb)
			ans = append(ans, append([]int(nil), comb...)) // slice copy
			return
		}
		if target < candidates[index] {
			return
		}
		for i := index; i < len(candidates); i++ {
			comb = append(comb, candidates[i])
			fmt.Printf("comb is %v\n", comb)
			dfs(candidates, target-candidates[i], i)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(candidates, target, 0)
	fmt.Printf("ans is %v\n", ans)
	return ans
}
