package greedy

import "testing"

func Test_maxProfitFee(t *testing.T) {
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
				prices: []int{1, 3, 2, 8},
				fee:    2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitFee(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfitFee() = %v, want %v", got, tt.want)
			}
		})
	}
}
