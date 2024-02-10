package binary_tree

import (
	"reflect"
	"testing"

	"algorithm/template"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *template.TreeNode
		p    *template.TreeNode
		q    *template.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *template.TreeNode
	}{
		// todo: 此测试用例有问题 指针指向问题
		//{
		//	name: "one",
		//	args: args{
		//		root: CreateBinaryTree([]int{3, 5, 1, 6, 2, 0, 8, Null, Null, 7, 4}),
		//		p:    CreateBinaryTree([]int{5, 6, 2, Null, Null, 7, 4}),
		//		q:    CreateBinaryTree([]int{1, 0, 8}),
		//	},
		//	want: CreateBinaryTree([]int{3, 5, 1, 6, 2, 0, 8, Null, Null, 7, 4}),
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
