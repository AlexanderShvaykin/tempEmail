package root

import (
	"github.com/AlexanderShvaykin/tempemail/pkg/cmd/gen"
	"github.com/AlexanderShvaykin/tempemail/pkg/cmd/mail"
	"github.com/AlexanderShvaykin/tempemail/pkg/cmdutil"
	"github.com/AlexanderShvaykin/tempemail/pkg/http"
	"github.com/spf13/cobra"
	"os"
)

func NewCmdRoot() *cobra.Command {
	factory := &cmdutil.Factory{
		Out:        os.Stdout,
		ErrOut:     os.Stderr,
		HttpClient: http.Client{},
	}

	cmd := &cobra.Command{
		Use:     "tmpemail",
		Short:   "Temp mailbox",
		Long:    "Create and check temp mailbox",
		Example: "tmpemail list",
	}

	cmd.AddCommand(gen.NewCmdGen(factory))
	cmd.AddCommand(mail.NewCmdList(factory))
	cmd.AddCommand(mail.NewCmdShow(factory))

	return cmd
}
