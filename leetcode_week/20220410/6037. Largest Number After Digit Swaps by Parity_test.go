package _0220410

import "testing"

func Test_largestInteger(t *testing.T) {
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
				num: 65875,
			},
			want: 87655,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestInteger(tt.args.num); got != tt.want {
				t.Errorf("largestInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
