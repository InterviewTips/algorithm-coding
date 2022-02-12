package binary_tree

import (
	"reflect"
	"testing"
)

func Test_largestValues(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "one",
			args: args{
				root: CreateBinaryTree([]int{1, 3, 2, 5, 3, Null, 9}),
			},
			want: []int{1, 3, 9},
		},
		{
			name: "two",
			args: args{
				root: CreateBinaryTree([]int{1, 2, 3}),
			},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestValues(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
