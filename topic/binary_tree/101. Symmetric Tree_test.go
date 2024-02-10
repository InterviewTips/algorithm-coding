package binary_tree

import (
	"testing"

	"algorithm/template"
)

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{
				root: template.CreateBinaryTree([]int{1, 2, 2, 3, 4, 4, 3}),
			},
			want: true,
		},
		{
			name: "two",
			args: args{
				root: template.CreateBinaryTree([]int{1, 2, 2, template.Null, 3, template.Null, 3}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
