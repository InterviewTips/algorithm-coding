package two_pointers

import (
	"reflect"
	"testing"

	"github.com/InterviewTips/algorithm-coding/guns"
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
		{
			name: "one",
			args: args{
				headA: guns.GenLinkList([]int{4, 1, 8, 4, 5}),
				headB: guns.GenLinkList([]int{5, 0, 1, 8, 4, 5}),
			},
			want: guns.GenLinkList([]int{8, 4, 5}),
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
