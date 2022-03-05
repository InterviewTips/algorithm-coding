package _0220305

import "testing"

func Test_findLUSlength(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				a: "aba",
				b: "cdc",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLUSlength(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("findLUSlength() = %v, want %v", got, tt.want)
			}
		})
	}
}
