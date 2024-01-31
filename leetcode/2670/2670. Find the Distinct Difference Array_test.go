package _670

import (
	"reflect"
	"testing"
)

func Test_distinctDifferenceArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "one",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: []int{-3, -1, 1, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distinctDifferenceArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distinctDifferenceArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
