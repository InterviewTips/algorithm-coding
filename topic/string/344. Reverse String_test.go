package string

import "testing"

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "one",
			args: args{
				s: []byte("hello"),
			},
		},
		{
			name: "two",
			args: args{
				s: []byte("Hannah"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
		})
	}
}
