package binary_tree

import (
	"reflect"
	"testing"

	"algorithm/template"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *template.TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "one",
			args: args{
				root:      template.CreateBinaryTree([]int{5, 4, 8, 11, template.Null, 13, 4, 7, 2, template.Null, template.Null, 5, 1}),
				targetSum: 22,
			},
			want: [][]int{
				{5, 4, 11, 2},
				{5, 8, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.targetSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
