package stack_queue

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "有效",
			args: args{s: "()[]{}"},
			want: true,
		},
		{
			name: "左边剩余",
			args: args{s: "(()[]{}"},
			want: false,
		},
		{
			name: "右边剩余",
			args: args{s: "()}"},
			want: false,
		},
		{
			name: "符号不等于",
			args: args{s: "(}"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
