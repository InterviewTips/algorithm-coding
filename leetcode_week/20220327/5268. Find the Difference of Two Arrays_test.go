package _0220327

import (
	"reflect"
	"testing"
)

func Test_findDifference(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "one",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{2, 3, 4},
			},
			want: [][]int{{1}, {4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDifference(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
