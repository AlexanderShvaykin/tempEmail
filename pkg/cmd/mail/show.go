package mail

import (
	"github.com/AlexanderShvaykin/tempemail/internal/secmail"
	"github.com/AlexanderShvaykin/tempemail/pkg/cmd/gen"
	"github.com/AlexanderShvaykin/tempemail/pkg/cmdutil"
	"github.com/jaytaylor/html2text"
	"github.com/spf13/cobra"
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
			text, err := html2text.FromString(mail.Body, html2text.Options{PrettyTables: false})
			if err != nil {
				panic(err)
			}
			cmdutil.Fprintf(
				f.Out,
				"From: %s. Date: %s\nSubject: %s\n%s\n",
				mail.From,
				mail.Date,
				mail.Subject,
				text,
			)
		},
	}
	cmd.Flags().StringVarP(&id, "id", "i", "", "Show an email by ID")
	_ = cmd.MarkFlagRequired("id")
	return cmd
}
