package _0220303

import "testing"

func Test_addDigits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				num: 38,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addDigits(tt.args.num); got != tt.want {
				t.Errorf("addDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
