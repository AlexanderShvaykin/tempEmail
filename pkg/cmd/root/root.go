package root

import (
	"github.com/spf13/cobra"
	"os"
	"tempEmail/pkg/cmd/gen"
	"tempEmail/pkg/cmd/mail"
	"tempEmail/pkg/cmdutil"
)

func NewCmdRoot() *cobra.Command {
	factory := cmdutil.Factory{
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}

	cmd := &cobra.Command{
		Use:     "tmpemail",
		Short:   "Temp mailbox",
		Long:    "Create and check temp mailbox",
		Example: "tmpemail list",
	}

	cmd.AddCommand(mail.NewCmdList())
	cmd.AddCommand(gen.NewCmdGen(&factory))

	return cmd
}
