package _0220225

import "testing"

func Test_complexNumberMultiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one",
			args: args{
				num1: "1+1i",
				num2: "1+1i",
			},
			want: "0+2i",
		},
		{
			name: "two",
			args: args{
				num1: "1+-1i",
				num2: "1+-1i",
			},
			want: "0+-2i",
		},
		{
			name: "err",
			args: args{
				num1: "1+-1i",
				num2: "0+0i",
			},
			want: "0+0i",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := complexNumberMultiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("complexNumberMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
