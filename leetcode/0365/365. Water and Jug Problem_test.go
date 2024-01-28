package _365

import "testing"

func Test_canMeasureWater(t *testing.T) {
	type args struct {
		jug1Capacity   int
		jug2Capacity   int
		targetCapacity int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "classic",
			args: args{
				jug1Capacity:   3,
				jug2Capacity:   5,
				targetCapacity: 4,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMeasureWater(tt.args.jug1Capacity, tt.args.jug2Capacity, tt.args.targetCapacity); got != tt.want {
				t.Errorf("canMeasureWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
