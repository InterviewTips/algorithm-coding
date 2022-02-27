package _0220227

import "sort"

func minimumTime(time []int, totalTrips int) int64 {
	sort.SliceStable(time, func(i, j int) bool {
		return time[i] < time[j]
	})
	lcmValue := time[0]
	for i := 1; i < len(time); i++ {
		lcmValue = lcm(time[i], lcmValue)
	}
	quanValue := 0
	for i := 0; i < len(time); i++ {
		quanValue += lcmValue / time[i]
	}
	quan := totalTrips / quanValue
	left := totalTrips % quanValue
	t := time[0]
	count := 0
	for {
		count = 0
		for i := 0; i < len(time); i++ {
			//log.Println("t is",t, "time is", time[i]
			count += t / time[i]
			if count >= left {
				goto end
			}
		}
		t += time[0]
	}
end:

	return int64(quan*lcmValue + t)
}

func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

func gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	} else {
		return y
	}
}
