package binary_tree

import (
	"reflect"
	"testing"

	"algorithm/template"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *template.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "one",
			args: args{
				root: template.CreateBinaryTree([]int{1, template.Null, 2, 3}),
			},
			want: []int{1, 3, 2},
		},
		{
			name: "two",
			args: args{
				root: template.CreateBinaryTree([]int{1, 2, 3, 4, 5}),
			},
			want: []int{4, 2, 5, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
