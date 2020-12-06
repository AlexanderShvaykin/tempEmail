package gen

import (
	"fmt"
	"github.com/spf13/cobra"
	"tempEmail/pkg/cmdutil"
	"tempEmail/pkg/util"
	"time"
)

const (
	nameLen = 10
)

func NewCmdGen(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate new email",
		Long:  "Generate new email or reset old email and save it to ENV",
		Run: func(c *cobra.Command, args []string) {
			userName := util.RandomString(nameLen, time.Now().UnixNano())
			tld := util.RandomTail(time.Now().UnixNano())
			domain := fmt.Sprintf("1secmail.%s", tld)
			email := Email{Login: userName, Domain: domain}
			_, err := fmt.Fprintln(f.Out, email.Email())
			if err != nil {
				panic(err)
			}
		},
	}

	return cmd
}

type Email struct {
	Login  string
	Domain string
}

func (e Email) Email() string {
	return fmt.Sprintf("%s@%s", e.Login, e.Domain)
}
