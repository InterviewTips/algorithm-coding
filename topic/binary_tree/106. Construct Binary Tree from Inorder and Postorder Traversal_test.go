package binary_tree

import (
	"reflect"
	"testing"
)

func Test_buildTreePostOrder(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "one",
			args: args{
				inorder:   []int{9, 3, 15, 20, 7},
				postorder: []int{9, 15, 7, 20, 3},
			},
			want: []int{3, 9, 20, Null, Null, 15, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildTreePostOrder(tt.args.inorder, tt.args.postorder)
			value := LevelOrderBinaryTreeAddNull(got)
			if !reflect.DeepEqual(value, tt.want) {
				t.Errorf("buildTreePostOrder() = %v, want %v", value, tt.want)
			}
		})
	}
}
