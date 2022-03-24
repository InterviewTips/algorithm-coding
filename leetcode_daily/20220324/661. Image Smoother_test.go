package _0220324

import (
	"reflect"
	"testing"
)

func Test_imageSmoother(t *testing.T) {
	type args struct {
		img [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "one",
			args: args{
				img: [][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := imageSmoother(tt.args.img); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imageSmoother() = %v, want %v", got, tt.want)
			}
		})
	}
}
