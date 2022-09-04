package _205

import "testing"

func Test_isIsomorphic(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{
				str1: "bbbaaaba",
				str2: "aaabbbab",
			},
			want: true,
		},
		{
			name: "one",
			args: args{
				str1: "bbbaaaab",
				str2: "aaabbbab",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
