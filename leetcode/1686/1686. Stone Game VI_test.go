package _686

import "testing"

func Test_stoneGameVI(t *testing.T) {
	type args struct {
		aliceValues []int
		bobValues   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				aliceValues: []int{1, 3},
				bobValues:   []int{2, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGameVI(tt.args.aliceValues, tt.args.bobValues); got != tt.want {
				t.Errorf("stoneGameVI() = %v, want %v", got, tt.want)
			}
		})
	}
}
