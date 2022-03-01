package theory

import "testing"

func Test_oneDimensionalCompletePack(t *testing.T) {
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
			want: 60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oneDimensionalCompletePack(tt.args.weight, tt.args.value, tt.args.bagWeight); got != tt.want {
				t.Errorf("oneDimensionalCompletePack() = %v, want %v", got, tt.want)
			}
		})
	}
}
