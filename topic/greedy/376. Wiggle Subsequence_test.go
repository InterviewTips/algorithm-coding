package greedy

import "testing"

func Test_wiggleMaxLength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				nums: []int{1, 7, 4, 9, 2, 5},
			},
			want: 6,
		},
		{
			name: "two",
			args: args{
				nums: []int{2, 2, 2},
			},
			want: 1, // 说明只能是 []int{2}
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wiggleMaxLength(tt.args.nums); got != tt.want {
				t.Errorf("wiggleMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
