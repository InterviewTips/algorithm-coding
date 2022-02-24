package greedy

import "testing"

func Test_canCompleteCircuit(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				gas:  []int{1, 2, 3, 4, 5},
				cost: []int{3, 4, 5, 1, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canCompleteCircuit(tt.args.gas, tt.args.cost); got != tt.want {
				t.Errorf("canCompleteCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_greedyCanCompleteCircuit(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				gas:  []int{1, 2, 3, 4, 5},
				cost: []int{3, 4, 5, 1, 2},
			},
			want: 3,
		},
		{
			name: "err",
			args: args{
				gas:  []int{4},
				cost: []int{5},
			},
			want: -1, // 在本题中 就是需要跑满一圈才可以的 所以这里是 -1
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := greedyCanCompleteCircuit(tt.args.gas, tt.args.cost); got != tt.want {
				t.Errorf("greedyCanCompleteCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}
