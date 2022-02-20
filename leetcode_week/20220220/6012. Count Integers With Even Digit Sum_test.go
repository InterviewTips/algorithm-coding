package _0220220

import "testing"

func Test_countEven(t *testing.T) {
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
				num: 30,
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countEven(tt.args.num); got != tt.want {
				t.Errorf("countEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
