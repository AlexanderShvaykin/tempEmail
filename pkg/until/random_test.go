package until

import "testing"

func TestRandomString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Returns random strings", args: args{n: 10}, want: RandomString(10)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomString(tt.args.n)
			if got == tt.want {
				t.Errorf("RandomString() return identity string %v", tt.want)
			}
			if len(got) != tt.args.n {
				t.Errorf("RandomString() return string with lenght not equal %v", tt.args.n)
			}
		})
	}
}
