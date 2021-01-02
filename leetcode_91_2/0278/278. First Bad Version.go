package _278

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// 找出 isBadVersion 为 true 的第一个元素 所以是靠左
func firstBadVersion(n int) int {
	left := 1
	right := n
	for left <= right {
		mid := left + (right-left)/2
		if isBadVersion(mid) { // true
			right = mid - 1
		} else { // false
			left = mid + 1
		}
	}
	return left
}

//mock
func isBadVersion(version int) bool {
	if version == 3 {
		return false
	}
	return true
}
