package binary_tree

import (
	"reflect"
	"testing"

	"algorithm/template"
)

func Test_constructMaximumBinaryTree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *template.TreeNode
	}{
		{
			name: "one",
			args: args{
				nums: []int{3, 2, 1, 6, 0, 5},
			},
			want: template.CreateBinaryTree([]int{6, 3, 5, template.Null, 2, 0, template.Null, template.Null, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructMaximumBinaryTree(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructMaximumBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
