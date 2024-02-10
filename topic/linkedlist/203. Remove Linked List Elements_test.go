package linkedlist

import (
	"reflect"
	"testing"

	"algorithm/template"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *template.ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *template.ListNode
	}{
		{
			name: "one",
			args: args{
				head: template.NewLinkedList([]int{1, 2, 6, 3, 4, 5, 6}),
				val:  6,
			},
			want: template.NewLinkedList([]int{1, 2, 3, 4, 5}),
		},
		{
			name: "two",
			args: args{
				head: template.NewLinkedList([]int{}),
				val:  1,
			},
			want: nil,
		},
		{
			name: "three",
			args: args{
				head: template.NewLinkedList([]int{7, 7, 7, 7}),
				val:  7,
			},
			want: template.NewLinkedList([]int{}),
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
