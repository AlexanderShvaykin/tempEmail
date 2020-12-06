package gen

import (
	"bytes"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strings"
	"tempEmail/pkg/cmdutil"
	"tempEmail/pkg/test"
	"testing"
)

func runCommand() (*test.CmdOut, error) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	factory := &cmdutil.Factory{
		Out: stdout, ErrOut: stderr,
	}

	cmd := NewCmdGen(factory)

	cmd.SetOut(ioutil.Discard)
	cmd.SetErr(ioutil.Discard)

	_, err := cmd.ExecuteC()
	return &test.CmdOut{
		OutBuf: stdout,
		ErrBuf: stderr,
	}, err
}

func TestGenCmd(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Generate new email and save to env"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := runCommand()
			if err != nil {
				t.Fatalf("error running command `generate`: %v", err)
			}

			r := regexp.MustCompile(`.{9}@1secmail\.(com|net|org)$`)
			email := strings.TrimSuffix(output.String(), "\n")
			if !r.MatchString(email) {
				t.Fatalf("output did not match regexp /%s/\n> output\n%q\n", r, email)
			}
			env := os.Getenv(EnvName)
			if env != email {
				t.Fatalf("Email did not save to env, current value %s, need %q", env, email)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want Email
	}{
		{
			name: "Splint address and returns new Email struct",
			args: args{address: "foo@baz.org"},
			want: Email{Login: "foo", Domain: "baz.org"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
