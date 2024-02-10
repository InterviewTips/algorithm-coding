package _0220306

import (
	"reflect"
	"testing"

	"algorithm/template"
)

func Test_createBinaryTree(t *testing.T) {
	type args struct {
		descriptions [][]int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "one",
			args: args{
				descriptions: [][]int{
					{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1},
				},
			},
			want: template.CreateBinaryTree([]int{50, 20, 80, 15, 17, 19}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createBinaryTree(tt.args.descriptions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
