package binary_tree

import (
	"reflect"
	"testing"

	"algorithm/template"
)

func Test_averageOfLevels(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "one",
			args: args{
				root: template.CreateBinaryTree([]int{3, 9, 20, template.Null, template.Null, 15, 7}),
			},
			want: []float64{3.00000, 14.50000, 11.00000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageOfLevels(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("averageOfLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}
