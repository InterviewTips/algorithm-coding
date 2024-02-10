package binary_tree

import (
	"testing"

	"algorithm/template"
)

func Test_minDepth(t *testing.T) {
	type args struct {
		root *template.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				root: template.CreateBinaryTree([]int{3, 9, 20, template.Null, template.Null, 15, 7}),
			},
			want: 2,
		},
		{
			name: "two",
			args: args{
				root: template.CreateBinaryTree([]int{2, template.Null, 3, template.Null, 4, template.Null, 5, template.Null, 6}),
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDepthLevel(t *testing.T) {
	type args struct {
		root *template.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				root: template.CreateBinaryTree([]int{3, 9, 20, template.Null, template.Null, 15, 7}),
			},
			want: 2,
		},
		{
			name: "two",
			args: args{
				root: template.CreateBinaryTree([]int{2, template.Null, 3, template.Null, 4, template.Null, 5, template.Null, 6}),
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepthLevel(tt.args.root); got != tt.want {
				t.Errorf("minDepthLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
