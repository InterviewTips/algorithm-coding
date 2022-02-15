package binary_tree

import (
	"reflect"
	"testing"
)

func Test_binaryTreePaths(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "one",
			args: args{
				root: CreateBinaryTree([]int{1, 2, 3, Null, 5}),
			},
			want: []string{"1->2->5", "1->3"},
		},
		{
			name: "two",
			args: args{
				root: CreateBinaryTree([]int{1}),
			},
			want: []string{"1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryTreePaths(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTreePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
