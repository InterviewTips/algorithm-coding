package binary_tree

import (
	"reflect"
	"testing"

	"algorithm/template"
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
			args: args{root: template.CreateNTree([]int{1, template.Null, 3, 2, 4, template.Null, 5, 6})},
			want: [][]int{
				{1},
				{3, 2, 4},
				{5, 6},
			},
		},
		{
			name: "two",
			args: args{
				root: template.CreateNTree(
					[]int{1, template.Null, 2, 3, 4, 5, template.Null, template.Null, 6, 7, template.Null, 8, template.Null, 9, 10, template.Null, template.Null, 11, template.Null, 12, template.Null, 13, template.Null, template.Null, 14},
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
