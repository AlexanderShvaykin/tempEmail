package main

import (
	"fmt"
	"github.com/AlexanderShvaykin/tempemail/pkg/cmd/root"
	"os"
)

func main() {
	cmd := root.NewCmdRoot()
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
