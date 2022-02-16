package binary_tree

import (
	"reflect"
	"testing"
)

func Test_mergeTrees(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "one",
			args: args{
				root1: CreateBinaryTree([]int{1, 3, 2, 5}),
				root2: CreateBinaryTree([]int{2, 1, 3, Null, 4, Null, 7}),
			},
			want: CreateBinaryTree([]int{3, 4, 5, 5, 4, Null, 7}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTrees(tt.args.root1, tt.args.root2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
