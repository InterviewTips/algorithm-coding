package dp

import "testing"

func Test_dpIsSubsequence(t *testing.T) {
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
				s: "abc",
				t: "ahbgdc",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dpIsSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("dpIsSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
