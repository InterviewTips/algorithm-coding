package binary_tree

import (
	"reflect"
	"testing"
)

func Test_nTreeLevelOrder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "one",
			args: args{root: CreateNTree([]int{1, Null, 3, 2, 4, Null, 5, 6})},
			want: [][]int{
				{1},
				{3, 2, 4},
				{5, 6},
			},
		},
		{
			name: "two",
			args: args{
				root: CreateNTree(
					[]int{1, Null, 2, 3, 4, 5, Null, Null, 6, 7, Null, 8, Null, 9, 10, Null, Null, 11, Null, 12, Null, 13, Null, Null, 14},
				),
			},
			want: [][]int{
				{1},
				{2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13},
				{14},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nTreeLevelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nTreeLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
