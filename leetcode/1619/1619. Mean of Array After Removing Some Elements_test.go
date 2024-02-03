package _619

import (
	"testing"
)

func Test_trimMean(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "one",
			args: args{
				[]int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3},
			},
			want: 2.00000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimMean(tt.args.arr); got != tt.want {
				t.Errorf("trimMean() = %v, want %v", got, tt.want)
			}
		})
	}
}
