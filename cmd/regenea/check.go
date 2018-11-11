package main

import (
	"gopkg.in/urfave/cli.v1"
)

func getCheckCommand() cli.Command {
	return cli.Command{
		Name:  "check",
		Usage: "check that the given genea file is valid (if no file is specified, stdin is used)",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "in",
				Usage: "input file to transform (default to stdin)",
			},
		},
		Action: doCheckCommand,
	}
}

func doCheckCommand(ctxt *cli.Context) error {
	infile := ""
	if ctxt.NArg() > 0 {
		infile = ctxt.Args()[0]
	}

	_, _, err := helperRead(infile, "genea")
	return err
}
