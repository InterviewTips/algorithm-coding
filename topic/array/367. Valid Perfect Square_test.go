package array

import "testing"

func Test_isPerfectSquare(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{16},
			want: true,
		},
		{
			name: "two",
			args: args{14},
			want: false,
		},
		{
			name: "0",
			args: args{0},
			want: true,
		},
		{
			name: "1",
			args: args{1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
