package cmdutil

import (
	"io"
	"tempEmail/pkg/http"
)

type Factory struct {
	Out        io.Writer
	ErrOut     io.Writer
	HttpClient http.Methods
}
