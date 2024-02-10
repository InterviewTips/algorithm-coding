package binary_tree

import (
	"reflect"
	"testing"

	"algorithm/template"
)

func Test_searchBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "one",
			args: args{
				root: template.CreateBinaryTree([]int{4, 2, 7, 1, 3}),
				val:  2,
			},
			want: template.CreateBinaryTree([]int{2, 1, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
