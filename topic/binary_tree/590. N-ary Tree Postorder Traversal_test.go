package binary_tree

import (
	"reflect"
	"testing"
)

func Test_postorder(t *testing.T) {
	type args struct {
		root *NTreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "one",
			args: args{
				root: CreateNTree([]int{1, Null, 3, 2, 4, Null, 5, 6}),
			},
			want: []int{5, 6, 3, 2, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postorderIteration(t *testing.T) {
	type args struct {
		root *NTreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "one",
			args: args{
				root: CreateNTree([]int{1, Null, 3, 2, 4, Null, 5, 6}),
			},
			want: []int{5, 6, 3, 2, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderIteration(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderIteration() = %v, want %v", got, tt.want)
			}
		})
	}
}
