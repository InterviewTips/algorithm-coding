package _165

import "testing"

func Test_compareVersion(t *testing.T) {
	type args struct {
		version1 string
		version2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				version1: "1.01",
				version2: "1.001",
			},
			want: 0,
		},
		{
			name: "two",
			args: args{
				version1: "1.0",
				version2: "1.0.0",
			},
			want: 0,
		},
		{
			name: "three",
			args: args{
				version1: "0.1",
				version2: "1.1",
			},
			want: -1,
		},
		{
			name: "err",
			args: args{
				version1: "1.0",
				version2: "1",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareVersion(tt.args.version1, tt.args.version2); got != tt.want {
				t.Errorf("compareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
