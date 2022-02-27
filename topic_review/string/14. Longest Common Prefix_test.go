package string

import (
	"log"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one",
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLetterByte(t *testing.T) {
	var letter byte
	log.Println(letter == 0)       // true
	log.Println(letter == byte(0)) // true
	log.Println(string(letter))
}
