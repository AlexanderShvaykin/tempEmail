package mail

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdList() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
		Run: func(c *cobra.Command, args []string) {
			fmt.Println(123)
		},
	}

	return cmd
}
