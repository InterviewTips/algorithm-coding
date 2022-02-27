package _0220227

import "testing"

func Test_minSteps(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				s: "leetcode",
				t: "coats",
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSteps(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
