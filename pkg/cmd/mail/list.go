package mail

import (
	"github.com/spf13/cobra"
	"strconv"
	"tempEmail/internal/secmail"
	"tempEmail/pkg/cmd/gen"
	"tempEmail/pkg/cmdutil"
)

func NewCmdList(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
		Run: func(c *cobra.Command, args []string) {
			email := gen.GenerateEmail()
			cmdutil.Fprintf(f.Out, "You email: %s\n", email.Email())
			list := secmail.GetMails(email.Login, email.Domain, f.HttpClient)
			if len(list) == 0 {
				cmdutil.Fprint(f.Out, "Mailbox is empty!")
			} else {
				for _, mail := range list {
					cmdutil.Fprintf(f.Out, "Mail ID: %s\n", strconv.FormatInt(mail.ID, 10))
				}
			}
		},
	}

	return cmd
}
