package binary_tree

import (
	"testing"

	"algorithm/template"
)

func Test_isSubtree(t *testing.T) {
	type args struct {
		root    *TreeNode
		subRoot *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{
				root:    template.CreateBinaryTree([]int{3, 4, 5, 1, 2}),
				subRoot: template.CreateBinaryTree([]int{4, 1, 2}),
			},
			want: true,
		},
		{
			name: "two",
			args: args{
				root: template.CreateBinaryTree([]int{
					1, template.Null, 1, template.Null, 1, template.Null, 1, template.Null, 1, template.Null, 1, template.Null, 1, template.Null, 1, template.Null, 1, template.Null, 1, template.Null, 1, 2,
				}),
				subRoot: template.CreateBinaryTree([]int{
					1, template.Null, 1, template.Null, 1, template.Null, 1, template.Null, 1, template.Null, 1, 2,
				}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.args.root, tt.args.subRoot); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
