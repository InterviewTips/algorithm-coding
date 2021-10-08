package linkedlist

import (
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "one",
			args: args{
				head: NewLinkedList([]int{1, 2, 6, 3, 4, 5, 6}),
				val:  6,
			},
			want: NewLinkedList([]int{1, 2, 3, 4, 5}),
		},
		{
			name: "two",
			args: args{
				head: NewLinkedList([]int{}),
				val:  1,
			},
			want: nil,
		},
		{
			name: "three",
			args: args{
				head: NewLinkedList([]int{7, 7, 7, 7}),
				val:  7,
			},
			want: NewLinkedList([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
