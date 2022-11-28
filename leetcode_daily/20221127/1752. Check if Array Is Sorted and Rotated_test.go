package _0221127

import "testing"

func Test_check(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{
				nums: []int{3, 4, 5, 1, 2},
			},
			want: true,
		},
		{
			name: "two",
			args: args{
				nums: []int{2, 1, 3, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.nums); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}
