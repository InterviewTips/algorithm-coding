package binary_tree

import (
	"reflect"
	"testing"

	"algorithm/template"
)

func Test_rightSideView(t *testing.T) {
	type args struct {
		root *template.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "one",
			args: args{
				root: template.CreateBinaryTree([]int{1, 2, 3, template.Null, 5, template.Null, 4}),
			},
			want: []int{1, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
