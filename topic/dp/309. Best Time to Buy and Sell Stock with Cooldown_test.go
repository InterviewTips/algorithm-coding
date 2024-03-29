package dp

import "testing"

func Test_maxProfit5(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				prices: []int{1, 2, 3, 0, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit5(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit5() = %v, want %v", got, tt.want)
			}
		})
	}
}
