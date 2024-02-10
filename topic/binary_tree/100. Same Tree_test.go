package binary_tree

import (
	"testing"

	"algorithm/template"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *template.TreeNode
		q *template.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{
				p: template.CreateBinaryTree([]int{1, 2, 3}),
				q: template.CreateBinaryTree([]int{1, 2, 3}),
			},
			want: true,
		},
		{
			name: "two",
			args: args{
				p: template.CreateBinaryTree([]int{1, 2}),
				q: template.CreateBinaryTree([]int{1, template.Null, 2}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
