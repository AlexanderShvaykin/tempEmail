package mail

import (
	"bytes"
	"io/ioutil"
	"regexp"
	"strings"
	"tempEmail/pkg/cmdutil"
	"tempEmail/pkg/httpstub"
	"tempEmail/pkg/test"
	"testing"
)

func runCommand(response string) (*test.CmdOut, error) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	factory := &cmdutil.Factory{
		Out: stdout, ErrOut: stderr, HttpClient: httpstub.HttpClient{Response: response},
	}

	cmd := NewCmdList(factory)

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
		name     string
		response string
		want     string
	}{
		{"Prints mail ids", `[{"id": 123}]`, "Mail ID: 123"},
		{"Print message about empty box", `[]`, "Mailbox is empty!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := runCommand(tt.response)
			if err != nil {
				t.Fatalf("error running command `list`: %v", err)
			}

			lines := strings.Split(output.String(), "\n")

			r := regexp.MustCompile(`You email: .*@1secmail\.(com|net|org)`)
			if !r.MatchString(lines[0]) {
				t.Fatalf("output did not match regexp /%s/\n> output\n%q\n", r, lines[0])
			}
			if lines[1] != tt.want {
				t.Fatalf("cmd output wrong, returs %v, want: %v", output.String(), tt.want)
			}
		})
	}
}
