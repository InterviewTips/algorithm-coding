package binary_tree

import "testing"

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{
				root:      CreateBinaryTree([]int{5, 4, 8, 11, Null, 13, 4, 7, 2, Null, Null, Null, 1}),
				targetSum: 22,
			},
			want: true,
		},
		{
			name: "two",
			args: args{
				root:      CreateBinaryTree([]int{1, 2, 3}),
				targetSum: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
