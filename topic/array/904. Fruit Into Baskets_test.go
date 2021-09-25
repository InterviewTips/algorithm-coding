package array

import "testing"

func Test_totalFruit(t *testing.T) {
	type args struct {
		fruits []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{fruits: []int{1, 2, 1}},
			want: 3,
		},
		{
			name: "two",
			args: args{fruits: []int{2, 3, 2, 2}},
			want: 4,
		},
		{
			name: "three",
			args: args{fruits: []int{1, 2, 3, 2, 2}},
			want: 4,
		},
		{
			name: "four",
			args: args{fruits: []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalFruit(tt.args.fruits); got != tt.want {
				t.Errorf("totalFruit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapLength(t *testing.T) {
	// map 删除 key 之后 len
	m := make(map[int]struct{})
	m[0] = struct{}{}
	m[1] = struct{}{}
	t.Log(len(m))
	delete(m, 0)
	t.Log(len(m))
	delete(m, 1)
	t.Log(len(m))
}
