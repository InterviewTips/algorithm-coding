package string

import (
	"reflect"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one",
			args: args{
				s: " hello   world",
			},
			want: "world hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseWords1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one",
			args: args{
				s: " hello   world ",
			},
			want: "world hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords1(tt.args.s); got != tt.want {
				t.Errorf("reverseWords1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeExtraSpace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "one",
			args: args{
				s: " hello   world ",
			},
			want: []byte("hello world"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeExtraSpace(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeExtraSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseSingleWord(t *testing.T) {
	type args struct {
		reverseBytes []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one",
			args: args{
				reverseBytes: []byte("olleh drow"),
			},
			want: "hello word",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseSingleWord(tt.args.reverseBytes); got != tt.want {
				t.Errorf("reverseSingleWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
