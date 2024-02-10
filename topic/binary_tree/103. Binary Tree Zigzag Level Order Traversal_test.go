package binary_tree

import (
	"reflect"
	"testing"

	"algorithm/template"
)

func Test_zigzagLevelOrder(t *testing.T) {
	type args struct {
		root *template.TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "one",
			args: args{
				root: template.CreateBinaryTree([]int{3, 9, 20, template.Null, template.Null, 15, 7}),
			},
			want: [][]int{
				{3}, {20, 9}, {15, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
