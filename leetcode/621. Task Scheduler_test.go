package leetcode

import "testing"

func Test_leastInterval(t *testing.T) {
	type args struct {
		tasks []byte
		n     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
				n:     2,
			},
			want: 8,
		},
		{
			name: "two",
			args: args{
				tasks: []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'},
				n:     2,
			},
			want: 16,
		},
		{
			name: "three",
			args: args{
				tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C', 'D', 'D', 'E'},
				n:     2,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastInterval(tt.args.tasks, tt.args.n); got != tt.want {
				t.Errorf("leastInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	a := make(map[int]*task)
	a[1] = &task{
		num: 1,
		n:   1,
	}
	t.Log(a[1])
	v, _ := a[1]
	v.num++

	t.Log(a[1])
}
