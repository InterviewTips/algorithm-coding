package leetcode

import "testing"

func Test_findShortestSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				nums: []int{1, 2, 2, 3, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findShortestSubArray(tt.args.nums); got != tt.want {
				t.Errorf("findShortestSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
