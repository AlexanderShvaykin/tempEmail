package root

import (
	"github.com/spf13/cobra"
	"tempEmail/pkg/cmd/mail"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "tmpemail",
		Short:   "Temp mailbox",
		Long:    "Create and check temp mailbox",
		Example: "tmpemail list",
	}

	cmd.AddCommand(mail.NewCmdList())

	return cmd
}
