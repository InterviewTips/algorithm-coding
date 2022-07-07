package _0220220

import "testing"

func Test_coutPairs(t *testing.T) {
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
				nums: []int{1, 2, 3, 4, 5},
				k:    2,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coutPairs(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("coutPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
