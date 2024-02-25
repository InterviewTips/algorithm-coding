package _694

import "testing"

func Test_reformatNumber(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one",
			args: args{number: "1-23-45 6"},
			want: "123-456",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reformatNumber(tt.args.number); got != tt.want {
				t.Errorf("reformatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
