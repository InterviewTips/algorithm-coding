package two_pointers

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
				s: []byte{'h', 'e', 'l', 'l', 'o'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			t.Log(string(tt.args.s))
		})
	}
}
