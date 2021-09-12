package array

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{4},
			want: 2,
		},
		{
			name: "two",
			args: args{8},
			want: 2,
		},
		{
			name: "0",
			args: args{0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mySqrt2(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{4},
			want: 2,
		},
		{
			name: "two",
			args: args{8},
			want: 2,
		},
		{
			name: "0",
			args: args{0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt2(tt.args.x); got != tt.want {
				t.Errorf("mySqrt2() = %v, want %v", got, tt.want)
			}
		})
	}
}
