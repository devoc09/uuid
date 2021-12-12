package main

import (
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

func (c *CLI) Run(args []string) int {
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
