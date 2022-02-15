package _0220215

import (
	"log"
	"reflect"
	"testing"
)

func Test_luckyNumbers(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "one",
			args: args{
				matrix: [][]int{
					{3, 7, 8}, {9, 11, 13}, {15, 16, 17},
				},
			},
			want: []int{15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := luckyNumbers(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("luckyNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareInts(t *testing.T) {
	a := [2]int{1}
	b := [2]int{1, 2}

	if a == b {
		log.Println("true")
	} else {
		log.Println("false")
	}
}
