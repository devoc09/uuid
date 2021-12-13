package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
)

const (
	ExitCodeOk             = 0
	ExitCodeParseFlagError = 1
	ExitCodeError          = 2
)

type CLI struct {
	outStream, errStream io.Writer
}

const usage = `uuid is a tool to generate uuid.
Usage: 
    %s [arguments]

Args:
    -h      Print Help message
    -v      Print the version of this tool
`

func (c *CLI) Run(args []string) int {
	var help, version bool

	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.Usage = func() {
		fmt.Fprintf(c.errStream, usage, Name)
	}
	flags.BoolVar(&help, "h", false, "display help message")
	flags.BoolVar(&version, "v", false, "display the version")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	if help {
		flags.Usage()
		return ExitCodeOk
	}

	if version {
		fmt.Fprintf(c.errStream, "%s v%s\n", Name, Version)
		return ExitCodeOk
	}

	id, err := uuid.NewRandom()
	if err != nil {
		fmt.Fprintf(c.errStream, "uuid generate Error\n")
		return ExitCodeError
	}

	fmt.Fprint(c.outStream, id.String())

	return ExitCodeOk
}

func main() {
	c := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(c.Run(os.Args))
}
