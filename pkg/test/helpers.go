package test

import (
	"bytes"
	"github.com/spf13/cobra"
	"io/ioutil"
	"tempEmail/pkg/cmdutil"
	"tempEmail/pkg/httpstub"
)

type CmdOut struct {
	OutBuf, ErrBuf *bytes.Buffer
}

func (c CmdOut) String() string {
	return c.OutBuf.String()
}

func (c CmdOut) Stderr() string {
	return c.ErrBuf.String()
}

func RunCommand(response string, args []string, builder func(*cmdutil.Factory) *cobra.Command) (*CmdOut, error) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	factory := &cmdutil.Factory{
		Out: stdout, ErrOut: stderr, HttpClient: &httpstub.HttpClient{Response: response},
	}

	cmd := builder(factory)

	cmd.SetOut(ioutil.Discard)
	cmd.SetErr(ioutil.Discard)
	cmd.SetArgs(args)

	_, err := cmd.ExecuteC()
	return &CmdOut{
		OutBuf: stdout,
		ErrBuf: stderr,
	}, err
}
