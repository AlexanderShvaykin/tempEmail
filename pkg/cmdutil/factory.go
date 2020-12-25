package cmdutil

import (
	"github.com/AlexanderShvaykin/tempEmail/pkg/http"
	"io"
)

type Factory struct {
	Out        io.Writer
	ErrOut     io.Writer
	HttpClient http.Methods
}
