package linkedlist

import (
	"github.com/InterviewTips/algorithm-coding/guns"
	"log"
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
				head: guns.GenLinkList([]int{1, 2, 6, 3, 4, 5, 6}),
				val:  6,
			},
			want: guns.GenLinkList([]int{1, 2, 3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElements(tt.args.head, tt.args.val)
			log.Printf("%p %p", tt.args.head, got) // 地址一致了
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
