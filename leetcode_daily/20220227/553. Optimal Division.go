package _0220227

import (
	"bytes"
	"fmt"
)

// todo: 看看 dp 解法
func optimalDivision(nums []int) string {
	if len(nums) == 1 {
		return fmt.Sprintf("%d", nums[0])
	}

	if len(nums) == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	}

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d/(%d", nums[0], nums[1]))
	for i := 2; i < len(nums)-1; i++ {
		buf.WriteString(fmt.Sprintf("/%d", nums[i]))
	}
	buf.WriteString(fmt.Sprintf("/%d)", nums[len(nums)-1]))

	return buf.String()
}
