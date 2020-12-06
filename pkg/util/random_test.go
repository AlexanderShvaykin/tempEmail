package util

import (
	"testing"
	"time"
)

func TestRandomString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Returns random strings", args: args{n: 10}, want: RandomString(10, time.Now().UnixNano())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomString(tt.args.n, time.Now().UnixNano())
			if got == tt.want {
				t.Errorf("RandomString() return identity string %v", tt.want)
			}
			if len(got) != tt.args.n {
				t.Errorf("RandomString() return string with lenght not equal %v", tt.args.n)
			}
		})
	}
}

func TestRandomTail(t *testing.T) {
	type args struct {
		seed int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Return some tail",
			args: args{1},
			want: "org",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomTail(tt.args.seed); got != tt.want {
				t.Errorf("RandomTail() = %v, want %v", got, tt.want)
			}
			if got := RandomTail(2); got == tt.want {
				t.Errorf("RandomTail() = %v, want %v", got, tt.want)
			}
		})
	}
}
