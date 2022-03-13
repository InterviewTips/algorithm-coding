package _0220313

import "testing"

func Test_maximumTop(t *testing.T) {
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
				nums: []int{5, 2, 2, 4, 0, 6},
				k:    4,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTop(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumTop() = %v, want %v", got, tt.want)
			}
		})
	}
}
