package binary_tree

import (
	"testing"

	"algorithm/template"
)

func Test_findTilt(t *testing.T) {
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
				root: template.CreateBinaryTree([]int{1, 2, 3}),
			},
			want: 1,
		},
		{
			name: "two",
			args: args{
				root: template.CreateBinaryTree([]int{4, 2, 9, 3, 5, template.Null, 7}),
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTilt(tt.args.root); got != tt.want {
				t.Errorf("findTilt() = %v, want %v", got, tt.want)
			}
		})
	}
}
