package greedy

import "testing"

func Test_findMinArrowShots(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				points: [][]int{
					{10, 16},
					{2, 8},
					{1, 6},
					{7, 12},
				},
			},
			want: 2,
		},
		{
			name: "two",
			args: args{
				points: [][]int{
					{1, 2},
					{3, 4},
					{5, 6},
					{7, 8},
				},
			},
			want: 4,
		},
		{
			name: "err",
			args: args{
				points: [][]int{
					{9, 12},
					{1, 10},
					{4, 11},
					{8, 12},
					{3, 9},
					{6, 9},
					{6, 7},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinArrowShots(tt.args.points); got != tt.want {
				t.Errorf("findMinArrowShots() = %v, want %v", got, tt.want)
			}
		})
	}
}
