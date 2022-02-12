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
				data: []int{1, Null, 2, 3},
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

func TestCreateNTree(t *testing.T) {
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
				data: []int{1, Null, 3, 2, 4, Null, 5, 6},
			},
		},
		{
			name: "two",
			args: args{
				data: []int{1, Null, 2, 3, 4, 5, Null, Null, 6, 7, Null, 8, Null, 9, 10, Null, Null, 11, Null, 12, Null, 13, Null, Null, 14},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateNTree(tt.args.data)
			LevelOrderLog(got)
		})
	}
}
