package binary_tree

import (
	"reflect"
	"testing"
)

func Test_connect(t *testing.T) {
	type args struct {
		root *PerfectNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "one",
			args: args{
				root: createPerfectTree([]int{1, 2, 3, 4, 5, 6, 7}),
			},
			want: []int{1, Null, 2, 3, Null, 4, 5, 6, 7, Null},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := connect(tt.args.root); !reflect.DeepEqual(levelOrderPerfectNode(got), tt.want) {
			//	t.Errorf("connect() = %v, want %v", levelOrderPerfectNode(got), tt.want)
			//}
			got := connect(tt.args.root)
			// 获取完美二叉树遍历结果
			res := levelOrderPerfectNode(got)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("connect() = %v, want %v", res, tt.want)
			}
		})
	}
}
