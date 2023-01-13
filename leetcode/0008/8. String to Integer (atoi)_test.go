package _008

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		numStr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{numStr: "42"},
			want: 42,
		},
		{
			name: "two",
			args: args{numStr: "   -42"},
			want: -42,
		},
		{
			name: "three",
			args: args{numStr: "4193 with words"},
			want: 4193,
		},
		{
			name: "45",
			args: args{numStr: "words and 987"},
			want: 0,
		},
		{
			name: "654",
			args: args{numStr: "-91283472332"},
			want: -2147483648,
		},
		{
			name: "764",
			args: args{numStr: "+1"},
			want: 1,
		},
		{
			name: "1076",
			args: args{numStr: "+-12"},
			want: 0,
		},
		{
			name: "1080",
			args: args{numStr: "00000-42a1234"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.numStr); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
