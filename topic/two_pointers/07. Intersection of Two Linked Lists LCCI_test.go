package two_pointers

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// todo: 此单测有问题
		//{
		//	name: "one",
		//	args: args{
		//		headA: guns.GenLinkList([]int{4, 1, 8, 4, 5}),
		//		headB: guns.GenLinkList([]int{5, 0, 1, 8, 4, 5}),
		//	},
		//	want: guns.GenLinkList([]int{8, 4, 5}),
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
