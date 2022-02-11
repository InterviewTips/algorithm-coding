package binary_tree

import (
	"testing"
)

func TestCreateBinaryTree(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "one",
			args: args{
				data: []int{1, null, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateBinaryTree(tt.args.data)
			PreOrder(got)
			InOrder(got)
			PostOrder(got)
		})
	}
}
