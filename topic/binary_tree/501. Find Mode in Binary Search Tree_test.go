package binary_tree

import (
	"reflect"
	"testing"

	"algorithm/template"
)

func Test_findMode(t *testing.T) {
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
				root: template.CreateBinaryTree([]int{1, template.Null, 2, 2}),
			},
			want: []int{2},
		},
		{
			name: "two",
			args: args{
				root: template.CreateBinaryTree([]int{2, 1, 5, template.Null, template.Null, 4, 5}),
			},
			want: []int{5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMode(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
