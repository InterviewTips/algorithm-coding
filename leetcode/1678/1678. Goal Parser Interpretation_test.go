package _678

import "testing"

func Test_interpret(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one",
			args: args{
				command: "G()(al)",
			},
			want: "Goal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := interpret(tt.args.command); got != tt.want {
				t.Errorf("interpret() = %v, want %v", got, tt.want)
			}
		})
	}
}
