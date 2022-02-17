package binary_tree

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{
				root: CreateBinaryTree([]int{2, 1, 3}),
			},
			want: true,
		},
		{
			name: "节点值不能重复",
			args: args{
				root: CreateBinaryTree([]int{2, 2, 2}),
			},
			want: false,
		},
		{
			name: "右子树没有满足需求",
			args: args{
				root: CreateBinaryTree([]int{5, 4, 6, Null, Null, 3, 7}),
			},
			want: false,
		},
		{
			name: "err",
			args: args{
				root: CreateBinaryTree([]int{120, 70, 140, 50, 100, 130, 160, 20, 55, 75, 110, 119, 135, 150, 200}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidBSTIteration(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "右子树没有满足需求",
			args: args{
				root: CreateBinaryTree([]int{5, 4, 6, Null, Null, 3, 7}),
			},
			want: false,
		},
		{
			name: "err",
			args: args{
				root: CreateBinaryTree([]int{120, 70, 140, 50, 100, 130, 160, 20, 55, 75, 110, 119, 135, 150, 200}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBSTIteration(tt.args.root); got != tt.want {
				t.Errorf("isValidBSTIteration() = %v, want %v", got, tt.want)
			}
		})
	}
}
