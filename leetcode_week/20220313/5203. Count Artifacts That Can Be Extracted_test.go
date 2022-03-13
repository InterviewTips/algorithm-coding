package _0220313

import "testing"

func Test_digArtifacts(t *testing.T) {
	type args struct {
		in0       int
		artifacts [][]int
		dig       [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				in0: 2,
				artifacts: [][]int{
					{0, 0, 0, 0},
					{0, 1, 1, 1},
				},
				dig: [][]int{
					{0, 0},
					{0, 1},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digArtifacts(tt.args.in0, tt.args.artifacts, tt.args.dig); got != tt.want {
				t.Errorf("digArtifacts() = %v, want %v", got, tt.want)
			}
		})
	}
}
