package _341

import (
	"reflect"
	"testing"
)

func Test_numberOfPairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "one",
			args: args{
				nums: []int{1, 3, 2, 1, 3, 2, 2},
			},
			want: []int{3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPairs(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numberOfPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
