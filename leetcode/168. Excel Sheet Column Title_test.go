package leetcode

import "testing"

func Test_convertToTitle(t *testing.T) {
	type args struct {
		columnNumber int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one",
			args: args{
				27,
			},
			want: "AA",
		},
		{
			name: "two",
			args: args{
				26,
			},
			want: "Z",
		},
		{
			name: "err",
			args: args{
				52,
			},
			want: "AZ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.args.columnNumber); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
