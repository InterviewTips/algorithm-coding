package binary_tree

import (
	"testing"

	"algorithm/template"
)

func Test_findBottomLeftValue(t *testing.T) {
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
				root: template.CreateBinaryTree([]int{2, 1, 3}),
			},
			want: 1,
		},
		{
			name: "two",
			args: args{
				root: template.CreateBinaryTree([]int{1, 2, 3, 4, template.Null, 5, 6, template.Null, template.Null, 7}),
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBottomLeftValue(tt.args.root); got != tt.want {
				t.Errorf("findBottomLeftValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
