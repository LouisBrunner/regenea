package main

import (
	"gopkg.in/urfave/cli.v1"
)

func getCheckCommand() cli.Command {
	return cli.Command{
		Name:  "check",
		Usage: "check that the given genea file is valid",
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
	_, err := helperRead(ctxt.String("in"), "genea")
	return err
}
