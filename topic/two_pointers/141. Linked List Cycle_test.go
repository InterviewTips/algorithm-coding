package two_pointers

import (
	"testing"

	"algorithm/template"
)

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	one := template.NewLinkedList([]int{3, 2, 0, -4})
	one.Next.Next.Next.Next = one.Next
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{
				head: one,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
