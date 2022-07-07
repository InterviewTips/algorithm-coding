package _0220306

import "testing"

func Test_minimalKSum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "one",
			args: args{
				nums: []int{5, 6},
				k:    6,
			},
			want: 25,
		},
		{
			name: "two",
			args: args{
				nums: []int{21, 63, 1, 98, 10, 77, 58, 10, 21, 99, 40, 27, 5, 67, 90, 29, 56, 16, 34, 25, 26, 17, 8, 3, 4, 17, 14, 32, 12, 17, 37, 81, 18, 49, 59, 23, 63, 39, 64, 7, 18, 35, 89, 11, 42, 30, 6, 15, 81, 52, 24, 39, 48, 9, 19, 34, 24, 2, 28, 13, 53, 4, 22, 20, 42, 2, 5, 58, 36, 31, 60, 38, 33},
				k:    20,
			},
			want: 1137,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimalKSum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minimalKSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
