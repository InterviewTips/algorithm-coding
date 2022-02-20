package backtracking

import (
	"reflect"
	"testing"
)

func Test_findSubsequences(t *testing.T) {
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
				nums: []int{4, 6, 7, 7},
			},
			want: [][]int{
				{4, 6},
				{4, 6, 7},
				{4, 6, 7, 7},
				{4, 7},
				{4, 7, 7},
				{6, 7},
				{6, 7, 7},
				{7, 7},
			},
		},
		{
			name: "two",
			args: args{
				nums: []int{4, 4, 3, 2, 1},
			},
			want: [][]int{
				{4, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubsequences(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubsequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
