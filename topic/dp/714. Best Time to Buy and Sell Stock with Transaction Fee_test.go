package dp

import "testing"

func Test_maxProfit6(t *testing.T) {
	type args struct {
		prices []int
		fee    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				prices: []int{1, 3, 2, 8, 4, 9},
				fee:    2,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit6(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfit6() = %v, want %v", got, tt.want)
			}
		})
	}
}
