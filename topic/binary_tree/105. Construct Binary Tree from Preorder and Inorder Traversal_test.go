package binary_tree

import (
	"reflect"
	"testing"

	"algorithm/template"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "one",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: []int{3, 9, 20, template.Null, template.Null, 15, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildTree(tt.args.preorder, tt.args.inorder)
			value := template.LevelOrderBinaryTreeAddNull(got)
			if !reflect.DeepEqual(value, tt.want) {
				t.Errorf("buildTree() = %v, want %v", value, tt.want)
			}
		})
	}
}
