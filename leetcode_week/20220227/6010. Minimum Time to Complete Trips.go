package _0220227

import "sort"

func minimumTime(time []int, totalTrips int) int64 {
	sort.SliceStable(time, func(i, j int) bool {
		return time[i] < time[j]
	})
	t := time[0]
	count := 0
	for {
		count = 0
		for i := 0; i < len(time); i++ {
			//log.Println("t is",t, "time is", time[i]
			count += t / time[i]
			if count >= totalTrips {
				goto end
			}
		}
		t += time[0]
	}
end:

	return int64(t)
}
