package binary_tree

import (
	"reflect"
	"testing"

	"algorithm/template"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "one",
			args: args{
				root: template.CreateBinaryTree([]int{4, 2, 7, 1, 3, 6, 9}),
			},
			want: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name: "two",
			args: args{
				root: template.CreateBinaryTree([]int{1, template.Null, 2, 3}),
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := invertTree(tt.args.root)
			value := template.LevelOrderBinaryTree(got)
			if !reflect.DeepEqual(value, tt.want) {
				t.Errorf("invertTree() = %v, want %v", value, tt.want)
			}
		})
	}
}
