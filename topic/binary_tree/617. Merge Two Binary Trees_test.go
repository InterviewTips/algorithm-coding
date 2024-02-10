package binary_tree

import (
	"reflect"
	"testing"

	"algorithm/template"
)

func Test_mergeTrees(t *testing.T) {
	type args struct {
		root1 *template.TreeNode
		root2 *template.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *template.TreeNode
	}{
		{
			name: "one",
			args: args{
				root1: template.CreateBinaryTree([]int{1, 3, 2, 5}),
				root2: template.CreateBinaryTree([]int{2, 1, 3, template.Null, 4, template.Null, 7}),
			},
			want: template.CreateBinaryTree([]int{3, 4, 5, 5, 4, template.Null, 7}),
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
