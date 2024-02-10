package template

import (
	"reflect"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "one",
			args: args{
				vals: []int{1, 2, 3},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLinkedList(tt.args.vals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}
