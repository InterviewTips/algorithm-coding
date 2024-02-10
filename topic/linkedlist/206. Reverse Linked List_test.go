package linkedlist

import (
	"reflect"
	"testing"

	"algorithm/template"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *template.ListNode
	}
	tests := []struct {
		name string
		args args
		want *template.ListNode
	}{
		{
			name: "one",
			args: args{
				head: template.NewLinkedList([]int{1, 2, 3, 4, 5}),
			},
			want: template.NewLinkedList([]int{5, 4, 3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
