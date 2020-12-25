package mail

import (
	"github.com/spf13/cobra"
	"tempEmail/internal/secmail"
	"tempEmail/pkg/cmd/gen"
	"tempEmail/pkg/cmdutil"
)

func NewCmdShow(f *cmdutil.Factory) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:     "show",
		Short:   "Print mail info",
		Long:    "Print mail info by ID",
		Example: "show c9k4y4pzr3@1secmail.com -i 123",
		Run: func(c *cobra.Command, args []string) {
			if len(args) == 0 {
				panic("Email is missing!")
			}
			email = gen.New(args[0])
			mail := secmail.GetMail(email.Login, email.Domain, id, f.HttpClient)
			//"From: batman@superhero.org. Date: 2018-06-08 14:33:55\nSubject: Super Man"
			cmdutil.Fprintf(
				f.Out,
				"From: %s. Date: %s\nSubject: %s\n%s",
				mail.From,
				mail.Date,
				mail.Subject,
				mail.Body,
			)
		},
	}
	cmd.Flags().StringVarP(&id, "id", "i", "", "Show an email by ID")
	_ = cmd.MarkFlagRequired("id")
	return cmd
}
