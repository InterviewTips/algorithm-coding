package _0220227

func minimumTime(time []int, totalTrips int) int64 {

	// 二分查找
	left := 1
	right := totalTrips * 1e7 // 上界

	// [left, right)
	for left < right {
		mid := left + (right-left)/2
		count := calcTrips(mid, time)
		if count < totalTrips {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return int64(left)
}

func calcTrips(t int, time []int) int {
	count := 0
	for i := 0; i < len(time); i++ {
		count += t / time[i]
	}

	return count
}
