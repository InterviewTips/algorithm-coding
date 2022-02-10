package stack_queue

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one",
			args: args{
				path: "/home/",
			},
			want: "/home",
		},
		{
			name: "two",
			args: args{
				path: "/../",
			},
			want: "/",
		},
		{
			name: "three",
			args: args{
				path: "/home//foo/",
			},
			want: "/home/foo",
		},
		{
			name: "four",
			args: args{
				path: "/a/./b/../../c/",
			},
			want: "/c",
		},
		{
			name: "err",
			args: args{
				path: "/a/../../b/../c//.//",
			},
			want: "/c",
		},
		{
			name: "err1",
			args: args{
				path: "/a//b////c/d//././/..",
			},
			want: "/a/b/c",
		},
		{
			name: "err2",
			args: args{
				path: "/..hidden",
			},
			want: "/..hidden",
		},
		{
			name: "err3",
			args: args{
				path: "/a//b////c/d//././/..",
			},
			want: "/a/b/c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
