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
			name:   "Returns email list",
			args:   args{login: "test_login", domain: "secmail.org"},
			want:   []Mail{{ID: int64(123)}},
			client: httpstub.HttpClient{Response: `[{"ID": 123}]`},
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
