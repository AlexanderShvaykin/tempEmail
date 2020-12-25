package mail

import (
	"regexp"
	"strings"
	"tempEmail/pkg/test"
	"testing"
)

func TestListCmd(t *testing.T) {
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
		args     []string
	}{
		{"Prints mail ids", mailsResponse, "Mail ID: 123. From: someone@example.com. Subject: Some subject. Date: 2018-06-08 14:33:55", []string{}},
		{"Print message about empty box", `[]`, "Mailbox is empty!", []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := test.RunCommand(tt.response, []string{""}, NewCmdList)
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
	//	with mail address
	t.Run("Check mails for user email", func(t *testing.T) {
		email := "foo@baz.ru"
		output, err := test.RunCommand(`[]`, []string{email}, NewCmdList)
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
