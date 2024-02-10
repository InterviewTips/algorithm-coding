package array

import "math"

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
// 注意：
//
// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
// 示例 1：
//
// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"
// 示例 2：
//
// 输入：s = "a", t = "a"
// 输出："a"
// 示例 3:
//
// 输入: s = "a", t = "aa"
// 输出: ""
// 解释: t 中两个字符 'a' 均应包含在 s 的子串中，
// 因此没有符合条件的子字符串，返回空字符串。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/minimum-window-substring
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func minWindow(source string, target string) string {
	// 两个 hash 表
	// 一个存 target 的
	// 一个存 source 的
	// 如何判断已经包含 可以写一个判断方法
	sourceMap := make(map[uint8]int)
	targetMap := make(map[uint8]int)
	// 初始化 targetMap
	initTargetMap(targetMap, target)
	// 开始计算
	result := math.MaxInt32
	indexL, indexR := 0, 0
	left := 0
	for right := 0; right < len(source); right++ {
		if targetMap[source[right]] > 0 {
			sourceMap[source[right]]++
		}
		for judgeOk(sourceMap, targetMap) && left <= right {
			subLength := right - left + 1
			if subLength < result {
				result = subLength
				indexL, indexR = left, left+result // right+1 最后返回用到
			}
			// 从 sourceMap 中去除一位 因为下面要滑动指针
			_, ok := targetMap[source[left]]
			if ok {
				sourceMap[source[left]]--
			}
			left++
		}

	}

	// 找不到结果返回 ""
	if result == math.MaxInt32 {
		return ""
	}

	return source[indexL:indexR]
}

func initTargetMap(targetMap map[uint8]int, target string) {
	for i := 0; i < len(target); i++ {
		//targetMap[target[i]] = targetMap[target[i]] + 1
		targetMap[target[i]]++
	}

	return
}

func judgeOk(source map[uint8]int, target map[uint8]int) bool {
	for k, v := range target {
		if source[k] < v {
			return false
		}
	}

	return true
}

// todo: 别人的解法
func minWindow1(s string, t string) string {
	win := make(map[byte]int)
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left := 0
	right := 0
	match := 0
	min := 1 << 30
	var start, end int
	for right < len(s) {
		tempC := s[right]
		right++
		if need[tempC] != 0 {
			win[tempC]++
		}
		if win[tempC] == need[tempC] && win[tempC] != 0 {
			match++
		}
		for match == len(need) {
			if min > right-left {
				min = right - left
				start = left
				end = right
			}
			tempC2 := s[left]
			left++
			if need[tempC2] != 0 {
				if need[tempC2] == win[tempC2] {
					match--
				}
				win[tempC2]--
			}
		}
	}
	if min == 1<<30 {
		return ""
	}
	return s[start:end]
}
