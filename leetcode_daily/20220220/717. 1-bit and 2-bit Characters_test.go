package _0220220

import "testing"

func Test_isOneBitCharacter(t *testing.T) {
	type args struct {
		bits []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{
				bits: []int{1, 0, 0},
			},
			want: true,
		},
		{
			name: "two",
			args: args{
				bits: []int{1, 1, 1, 0},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOneBitCharacter(tt.args.bits); got != tt.want {
				t.Errorf("isOneBitCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
