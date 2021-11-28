package hashtable

import (
	"log"
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "one",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "two",
			args: args{
				nums: []int{},
			},
			want: [][]int{},
		},
		{
			name: "zero",
			args: args{
				nums: []int{0, 0, 0, 0, 0, 0, 0},
			},
			want: [][]int{
				{0, 0, 0},
			},
		},
		{
			name: "nil",
			args: args{
				nums: []int{0},
			},
			want: [][]int{},
		},
		{
			name: "w",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4},
			},
			want: [][]int{
				{-4, 0, 4},
				{-4, 1, 3},
				{-3, -1, 4},
				{-3, 0, 3},
				{-3, 1, 2},
				{-2, -1, 3},
				{-2, 0, 2},
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpecial(t *testing.T) {
	log.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}))
}
