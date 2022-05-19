package _0220519

import "testing"

func Test_minMoves2(t *testing.T) {
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
				nums: []int{1, 2, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMoves2(tt.args.nums); got != tt.want {
				t.Errorf("minMoves2() = %v, want %v", got, tt.want)
			}
		})
	}
}
