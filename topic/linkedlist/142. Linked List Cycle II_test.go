package linkedlist

import (
	"reflect"
	"testing"

	"algorithm/template"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	linkedList := template.NewLinkedList([]int{3, 2, 0, -4})
	linkedList.Next.Next.Next.Next = linkedList.Next // -4 -> 2
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "one",
			args: args{
				head: linkedList,
			},
			want: linkedList.Next,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
