package binary_tree

import (
	"testing"

	"algorithm/template"
)

func Test_getMinimumDifference(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				root: template.CreateBinaryTree([]int{4, 2, 6, 1, 3}),
			},
			want: 1,
		},
		{
			name: "err",
			args: args{
				root: template.CreateBinaryTree([]int{1, 0, 48, template.Null, template.Null, 12, 49}),
			},
			want: 1,
		},
		{
			name: "err1",
			args: args{
				root: template.CreateBinaryTree([]int{5, 4, 7}),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinimumDifference(tt.args.root); got != tt.want {
				t.Errorf("getMinimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
