package mail

import (
	"tempEmail/pkg/test"
	"testing"
)

func TestShowCmd(t *testing.T) {
	mailResponse, err := test.Fixture("mail.json")
	if err != nil {
		t.Fatal("read file Error!")
	}
	tests := []struct {
		name     string
		response string
		want     string
		args     []string
	}{
		{
			"Prints mail ids",
			mailResponse,
			"From: batman@superhero.org. Date: 2018-06-08 14:33:55\nSubject: Super Man\nSome message body\n",
			[]string{"foo@baz.org", "-i 123"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := test.RunCommand(tt.response, tt.args, NewCmdShow)
			if err != nil {
				t.Fatalf("error running command `list`: %v", err)
			}
			expected := output.String()
			if expected != tt.want {
				t.Errorf("It doesn't return mail info, returned: %v, want: %v", expected, tt.want)
			}
		})
	}
}
