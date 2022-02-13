package _0220213

import "testing"

func Test_countOperations(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				num1: 2,
				num2: 3,
			},
			want: 3,
		},
		{
			name: "two",
			args: args{
				num1: 10,
				num2: 10,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOperations(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("countOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
