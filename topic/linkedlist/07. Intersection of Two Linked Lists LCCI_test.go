package linkedlist

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
		// todo: 这个测试用例不能这么写 会无法通过 指针指向不同
		{
			name: "one",
			args: args{
				headA: NewLinkedList([]int{4, 1, 8, 4, 5}),
				headB: NewLinkedList([]int{5, 0, 1, 8, 4, 5}),
			},
			want: NewLinkedList([]int{8, 4, 5}),
		},
		{
			name: "two",
			args: args{
				headA: NewLinkedList([]int{1, 2}),
				headB: NewLinkedList([]int{3, 4}),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
