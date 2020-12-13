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

func runCommand(response string, args []string) (*test.CmdOut, error) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	factory := &cmdutil.Factory{
		Out: stdout, ErrOut: stderr, HttpClient: &httpstub.HttpClient{Response: response},
	}

	cmd := NewCmdList(factory)

	cmd.SetOut(ioutil.Discard)
	cmd.SetErr(ioutil.Discard)
	cmd.SetArgs(args)

	_, err := cmd.ExecuteC()
	return &test.CmdOut{
		OutBuf: stdout,
		ErrBuf: stderr,
	}, err
}

func TestGenCmd(t *testing.T) {
	mailsResponse := `
[{
	"id": 123,
	"from": "someone@example.com",
	"subject": "Some subject",
	"date": "2018-06-08 14:33:55"
}]
`

	tests := []struct {
		name     string
		response string
		want     string
	}{
		{"Prints mail ids", mailsResponse, "Mail ID: 123. From: someone@example.com. Subject: Some subject. Date: 2018-06-08 14:33:55"},
		{"Print message about empty box", `[]`, "Mailbox is empty!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := runCommand(tt.response, []string{""})
			if err != nil {
				t.Fatalf("error running command `list`: %v", err)
			}

			lines := strings.Split(output.String(), "\n")

			r := regexp.MustCompile(`You email: .*@1secmail\.(com|net|org)`)
			if !r.MatchString(lines[0]) {
				t.Fatalf("output did not match regexp /%s/\n> output\n%q\n", r, lines[0])
			}
			if lines[1] != tt.want {
				t.Fatalf("cmd output wrong, returs %v, want: %v", lines[1], tt.want)
			}
		})
	}
	//	whit mail address
	t.Run("Check mails for user email", func(t *testing.T) {
		email := "foo@baz.ru"
		output, err := runCommand(`[]`, []string{email})
		if err != nil {
			t.Fatalf("error running command `list`: %v", err)
		}

		lines := strings.Split(output.String(), "\n")
		want := "You email: " + email
		if lines[0] != "You email: "+email {
			t.Fatalf("cmd output wrong, returs %v, want: %v", lines[0], want)
		}
	})
}
