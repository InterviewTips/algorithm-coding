package _1927

import "testing"

func Test_sumGame(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{
				num: "5023",
			},
			want: false,
		},
		{
			name: "two",
			args: args{
				num: "?3295???",
			},
			want: false,
		},
		{
			name: "three",
			args: args{
				num: "25??",
			},
			want: true,
		},
		{
			name: "err",
			args: args{
				num: "2582?9",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumGame(tt.args.num); got != tt.want {
				t.Errorf("sumGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
