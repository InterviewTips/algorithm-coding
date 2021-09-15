package array

import "testing"

func Test_backspaceCompare1(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{
				s: "ab#c",
				t: "ad#c",
			},
			want: true,
		},
		{
			name: "two",
			args: args{
				s: "ab##",
				t: "c#d#",
			},
			want: true,
		},
		{
			name: "three",
			args: args{
				s: "a##c",
				t: "#a#c",
			},
			want: true,
		},
		{
			name: "four",
			args: args{
				s: "a#c",
				t: "b",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare1(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("backspaceCompare1() = %v, want %v", got, tt.want)
			}
		})
	}
}
