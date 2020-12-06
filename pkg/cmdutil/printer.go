package cmdutil

import (
	"fmt"
	"io"
)

func Fprintf(w io.Writer, format string, a ...interface{}) {
	_, err := fmt.Fprintf(w, format, a...)
	if err != nil {
		panic(err)
	}
}

func Fprint(w io.Writer, a ...interface{}) {
	_, err := fmt.Fprint(w, a...)
	if err != nil {
		panic(err)
	}
}
