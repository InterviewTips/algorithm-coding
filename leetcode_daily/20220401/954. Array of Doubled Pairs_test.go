package _0220401

import "testing"

func Test_canReorderDoubled(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{
				arr: []int{2, 4, 0, 0, 8, 1},
			},
			want: true,
		},
		{
			name: "two",
			args: args{
				arr: []int{1, 1, 2, 2},
			},
			want: true,
		},
		{
			name: "three",
			args: args{
				arr: []int{2, 3, 4, 6},
			},
			want: true,
		},
		{
			name: "err",
			args: args{
				arr: []int{-2, -6, -3, 4, -4, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canReorderDoubled(tt.args.arr); got != tt.want {
				t.Errorf("canReorderDoubled() = %v, want %v", got, tt.want)
			}
		})
	}
}
