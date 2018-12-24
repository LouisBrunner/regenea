package main

import (
	"gopkg.in/urfave/cli.v1"
)

func getDisplayCommand() cli.Command {
	return cli.Command{
		Name: "display",
		// Usage: "check that the given genea file is valid (if no file is specified, stdin is used)",
		// Flags: []cli.Flag{
		// 	inField,
		// 	subsetField,
		// },
		Action: doDisplayCommand,
	}
}

func doDisplayCommand(ctxt *cli.Context) error {
	panic("not implemented")
}
