package stack_queue

import "testing"

func TestMyQueue_Pop(t *testing.T) {
	type fields struct {
		Stack1 []int
		Stack2 []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "one",
			fields: fields{
				Stack1: []int{1, 2},
				Stack2: []int{},
			},
			want: 1,
		},
		{
			name: "two",
			fields: fields{
				Stack1: []int{},
				Stack2: []int{2, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MyQueue{
				Stack1: tt.fields.Stack1,
				Stack2: tt.fields.Stack2,
			}
			if got := m.Pop(); got != tt.want {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor(t *testing.T) {
	m := Constructor()
	m.Push(1)
	m.Push(2)
	t.Log("pop is", m.Pop())
	m.Push(3)
	m.Push(4)
	t.Log("peek", m.Peek())
	t.Log("peek", m.Peek())
}

func TestCase(t *testing.T) {
	m := Constructor()
	m.Push(1)
	t.Log("pop is", m.Pop())
	t.Log("empty", m.Empty())
}
