package binary_tree

import "testing"

func Test_verifySequenceOfBST(t *testing.T) {
	type args struct {
		sequence []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{
				sequence: []int{1, 3, 2, 6, 5},
			},
			want: true,
		},
		{
			name: "two",
			args: args{
				sequence: []int{1, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifySequenceOfBST(tt.args.sequence); got != tt.want {
				t.Errorf("verifySequenceOfBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
