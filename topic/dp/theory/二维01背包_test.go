package theory

import "testing"

func Test_twoDimensional01Bag(t *testing.T) {
	type args struct {
		weight    []int
		value     []int
		bagWeight int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				weight:    []int{1, 3, 4},
				value:     []int{15, 20, 30},
				bagWeight: 4,
			},
			want: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoDimensional01Bag(tt.args.weight, tt.args.value, tt.args.bagWeight); got != tt.want {
				t.Errorf("twoDimensional01Bag() = %v, want %v", got, tt.want)
			}
		})
	}
}
