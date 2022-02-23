package greedy

import "testing"

func Test_largestSumAfterKNegations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				nums: []int{4, 2, 3},
				k:    1,
			},
			want: 5,
		},
		{
			name: "two",
			args: args{
				nums: []int{3, -1, 0, 2},
				k:    3,
			},
			want: 6,
		},
		{
			name: "three",
			args: args{
				nums: []int{8, -7, -3, -9, 1, 9, -6, -9, 3},
				k:    8,
			},
			want: 53,
		},
		{
			name: "err",
			args: args{
				nums: []int{-4, -2, -3},
				k:    4,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSumAfterKNegations(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("largestSumAfterKNegations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_absLargestSumAfterKNegations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				nums: []int{4, 2, 3},
				k:    1,
			},
			want: 5,
		},
		{
			name: "two",
			args: args{
				nums: []int{3, -1, 0, 2},
				k:    3,
			},
			want: 6,
		},
		{
			name: "three",
			args: args{
				nums: []int{8, -7, -3, -9, 1, 9, -6, -9, 3},
				k:    8,
			},
			want: 53,
		},
		{
			name: "err",
			args: args{
				nums: []int{-4, -2, -3},
				k:    4,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := absLargestSumAfterKNegations(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("absLargestSumAfterKNegations() = %v, want %v", got, tt.want)
			}
		})
	}
}
