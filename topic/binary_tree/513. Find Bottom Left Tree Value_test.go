package binary_tree

import "testing"

func Test_findBottomLeftValue(t *testing.T) {
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
				root: CreateBinaryTree([]int{2, 1, 3}),
			},
			want: 1,
		},
		{
			name: "two",
			args: args{
				root: CreateBinaryTree([]int{1, 2, 3, 4, Null, 5, 6, Null, Null, 7}),
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBottomLeftValue(tt.args.root); got != tt.want {
				t.Errorf("findBottomLeftValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
