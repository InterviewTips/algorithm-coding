package _0220319

import (
	"testing"

	"algorithm/template"
)

func Test_tree2str(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one",
			args: args{
				root: template.CreateBinaryTree([]int{1, 2, 3, 4}),
			},
			want: "1(2(4))(3)",
		},
		{
			name: "two",
			args: args{
				root: template.CreateBinaryTree([]int{1, 2, 3, template.Null, 4}),
			},
			want: "1(2()(4))(3)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree2str(tt.args.root); got != tt.want {
				t.Errorf("tree2str() = %v, want %v", got, tt.want)
			}
		})
	}
}
