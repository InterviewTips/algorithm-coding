package dp

import "testing"

func Test_minDistance2(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				word1: "horse",
				word2: "ros",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance2(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance2() = %v, want %v", got, tt.want)
			}
		})
	}
}
