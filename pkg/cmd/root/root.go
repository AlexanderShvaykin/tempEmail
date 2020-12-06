package root

import (
	"github.com/spf13/cobra"
	"os"
	"tempEmail/pkg/cmd/gen"
	"tempEmail/pkg/cmd/mail"
	"tempEmail/pkg/cmdutil"
	"tempEmail/pkg/http"
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

	cmd.AddCommand(mail.NewCmdList(factory))
	cmd.AddCommand(gen.NewCmdGen(factory))

	return cmd
}
