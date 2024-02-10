package greedy

import (
	"testing"

	"algorithm/template"
)

func Test_minCameraCover(t *testing.T) {
	type args struct {
		root *template.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				root: template.CreateBinaryTree([]int{0, 0, template.Null, 0, 0}),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCameraCover(tt.args.root); got != tt.want {
				t.Errorf("minCameraCover() = %v, want %v", got, tt.want)
			}
		})
	}
}
