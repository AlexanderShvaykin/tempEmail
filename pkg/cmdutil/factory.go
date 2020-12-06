package cmdutil

import "io"

type Factory struct {
	Out    io.Writer
	ErrOut io.Writer
}
