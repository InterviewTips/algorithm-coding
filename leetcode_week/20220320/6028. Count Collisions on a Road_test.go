package _0220320

import "testing"

func Test_countCollisions(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				data: "RLRSLL",
			},
			want: 5,
		},
		{
			name: "two",
			args: args{
				data: "SSRSSRLLRSLLRSRSSRLRRRRLLRRLSSRR",
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCollisions(tt.args.data); got != tt.want {
				t.Errorf("countCollisions() = %v, want %v", got, tt.want)
			}
		})
	}
}
