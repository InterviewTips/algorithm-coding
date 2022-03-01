package dp

import "testing"

func Test_climbStairsCompletePack(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				n: 2,
				m: 2,
			},
			want: 2,
		},
		{
			name: "two",
			args: args{
				n: 3,
				m: 2,
			},
			want: 3,
		},
		{
			name: "three",
			args: args{
				n: 3,
				m: 3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairsCompletePack(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("climbStairsCompletePack() = %v, want %v", got, tt.want)
			}
		})
	}
}
