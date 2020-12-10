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

func Fprintln(w io.Writer, a ...interface{}) {
	_, err := fmt.Fprintln(w, a...)
	if err != nil {
		panic(err)
	}
}
