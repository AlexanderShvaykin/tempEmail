package secmail

import (
	"github.com/AlexanderShvaykin/tempemail/pkg/httpstub"
	"github.com/AlexanderShvaykin/tempemail/pkg/test"
	"reflect"
	"testing"
)

func TestGetMails(t *testing.T) {
	type args struct {
		login  string
		domain string
	}
	tests := []struct {
		name   string
		args   args
		want   []Mail
		client httpstub.HttpClient
	}{
		{
			name: "Returns email list",
			args: args{login: "test_login", domain: "secmail.org"},
			want: []Mail{
				{ID: int64(98449250), From: "batman@superhero.org", Subject: "Super Man", Date: "2020-12-06 16:37:08"},
			},
			client: httpstub.HttpClient{Response: `[{"id":98449250,"from":"batman@superhero.org","subject":"Super Man","date":"2020-12-06 16:37:08"}]`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMails(tt.args.login, tt.args.domain, &tt.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMail(t *testing.T) {
	mailResponse, err := test.Fixture("mail.json")
	if err != nil {
		t.Fatal("read file Error!")
	}
	type args struct {
		login  string
		domain string
		id     string
	}
	tests := []struct {
		name   string
		args   args
		want   Mail
		client httpstub.HttpClient
	}{
		{
			name: "Returns email list",
			args: args{login: "test_login", domain: "secmail.org", id: "123"},
			want: Mail{
				ID: int64(639), From: "batman@superhero.org", Subject: "Super Man", Date: "2018-06-08 14:33:55", Body: "Some message body\n\n",
			},
			client: httpstub.HttpClient{Response: mailResponse},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMail(tt.args.login, tt.args.domain, tt.args.id, &tt.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMails() = %v, want %v", got, tt.want)
			}
			if baseUrl != tt.client.Url {
				t.Errorf("Dont send request to %v, sends to %v", baseUrl, tt.client.Url)
			}
			params := map[string]string{
				"action": "readMessage",
				"login":  tt.args.login,
				"domain": tt.args.domain,
				"id":     tt.args.id,
			}
			if !reflect.DeepEqual(params, tt.client.Args) {
				t.Errorf("Dont send request with args %v, sends with %v", params, tt.client.Args)
			}
		})
	}
}
