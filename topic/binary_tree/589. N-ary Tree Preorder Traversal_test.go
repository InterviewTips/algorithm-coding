package binary_tree

import (
	"reflect"
	"testing"
)

func Test_preorder(t *testing.T) {
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
			want: []int{1, 3, 5, 6, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preorderIteration(t *testing.T) {
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
			want: []int{1, 3, 5, 6, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderIteration(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderIteration() = %v, want %v", got, tt.want)
			}
		})
	}
}
