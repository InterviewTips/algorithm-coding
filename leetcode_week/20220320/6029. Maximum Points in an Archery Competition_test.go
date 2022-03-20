package _0220320

import (
	"reflect"
	"testing"
)

func Test_maximumBobPoints(t *testing.T) {
	type args struct {
		numArrows   int
		aliceArrows []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "one",
			args: args{
				numArrows:   9,
				aliceArrows: []int{1, 1, 0, 1, 0, 0, 2, 1, 0, 1, 2, 0},
			},
			want: []int{0, 0, 0, 0, 1, 1, 0, 0, 1, 2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumBobPoints(tt.args.numArrows, tt.args.aliceArrows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumBobPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
