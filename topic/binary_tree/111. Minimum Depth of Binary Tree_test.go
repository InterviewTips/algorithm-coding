package binary_tree

import "testing"

func Test_minDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				root: CreateBinaryTree([]int{3, 9, 20, Null, Null, 15, 7}),
			},
			want: 2,
		},
		{
			name: "two",
			args: args{
				root: CreateBinaryTree([]int{2, Null, 3, Null, 4, Null, 5, Null, 6}),
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDepthLevel(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				root: CreateBinaryTree([]int{3, 9, 20, Null, Null, 15, 7}),
			},
			want: 2,
		},
		{
			name: "two",
			args: args{
				root: CreateBinaryTree([]int{2, Null, 3, Null, 4, Null, 5, Null, 6}),
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepthLevel(tt.args.root); got != tt.want {
				t.Errorf("minDepthLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
