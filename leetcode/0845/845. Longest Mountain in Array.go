package _845

func longestMountain(A []int) int {
	n := len(A)
	left := 0
	ans := 0
	for left+2 < n {
		right := left + 1
		if A[left] < A[left+1] {
			for right+1 < n && A[right] < A[right+1] {
				right++
			}
			if right < n-1 && A[right] > A[right+1] { // 到达顶峰
				for right+1 < n && A[right] > A[right+1] {
					right++
				}
				if right-left+1 > ans {
					ans = right - left + 1
				}
			} else { // right + 1 为下一个左侧山脚
				right = right + 1
			}
		}
		left = right
	}
	return ans
}
