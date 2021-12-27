package string

import "testing"

func Test_reverseLeftWords(t *testing.T) {
	type args struct {
		source string
		num    int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one",
			args: args{
				source: "abcdefg",
				num:    2,
			},
			want: "cdefgab",
		},
		{
			name: "two",
			args: args{
				source: "lrloseumgh",
				num:    6,
			},
			want: "umghlrlose",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLeftWords(tt.args.source, tt.args.num); got != tt.want {
				t.Errorf("reverseLeftWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
