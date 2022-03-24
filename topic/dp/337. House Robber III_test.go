package dp

import (
	"testing"

	"algorithm/topic/binary_tree"
)

var Null = binary_tree.Null

func Test_rob3(t *testing.T) {
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
				root: binary_tree.CreateBinaryTree([]int{3, 2, 3, Null, 3, Null, 1}),
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob3(tt.args.root); got != tt.want {
				t.Errorf("rob3() = %v, want %v", got, tt.want)
			}
		})
	}
}
