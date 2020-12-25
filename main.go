package main

import (
	"fmt"
	"os"
	"tempEmail/pkg/cmd/root"
)

func main() {
	cmd := root.NewCmdRoot()
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
