package secmail

import (
	"reflect"
	"tempEmail/pkg/httpstub"
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
			if got := GetMails(tt.args.login, tt.args.domain, tt.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMails() = %v, want %v", got, tt.want)
			}
		})
	}
}
