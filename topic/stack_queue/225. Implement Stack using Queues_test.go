package stack_queue

import "testing"

func TestMyStack_Pop(t *testing.T) {
	type fields struct {
		Queue1 []int
		Queue2 []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "one",
			fields: fields{
				Queue1: []int{1},
				Queue2: []int{2, 3},
			},
			want: 3,
		},
		{
			name: "two",
			fields: fields{
				Queue1: []int{1, 2},
				Queue2: []int{},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MyStack{
				Queue1: tt.fields.Queue1,
				Queue2: tt.fields.Queue2,
			}
			if got := m.Pop(); got != tt.want {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
