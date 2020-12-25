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
		Use:     "list",
		Short:   "Print list of mails",
		Long:    "Print list of mails and generate email for you, also you can add emails params, for example, list myemail@1sec.com",
		Example: "list c9k4y4pzr3@1secmail.com",
		Run: func(c *cobra.Command, args []string) {
			if len(args) > 0 {
				currentEmail = args[0]
			}
			if currentEmail == "" {
				email = gen.GenerateEmail()
			} else {
				email = gen.New(currentEmail)
			}
			cmdutil.Fprintf(f.Out, "You email: %s\n", email.Email())
			list := secmail.GetMails(email.Login, email.Domain, f.HttpClient)
			if len(list) == 0 {
				cmdutil.Fprintln(f.Out, "Mailbox is empty!")
			} else {
				for _, mail := range list {
					cmdutil.Fprintf(
						f.Out,
						"Mail ID: %s. From: %s. Subject: %s. Date: %s\n",
						strconv.FormatInt(mail.ID, 10),
						mail.From,
						mail.Subject,
						mail.Date,
					)
				}
			}
		},
	}
	return cmd
}
