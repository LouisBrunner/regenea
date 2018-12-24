package main

import (
	"fmt"

	"gopkg.in/urfave/cli.v1"
)

func getQueryCommand() cli.Command {
	return cli.Command{
		Name:  "query",
		Usage: "query information from the given genea file (if no file is specified, stdin is used)",
		Flags: []cli.Flag{
			inField,
			subsetField,
			cli.StringFlag{
				Name:  "id",
				Usage: "show the person with the given id",
			},
		},
		Action: doQueryCommand,
	}
}

func doQueryCommand(ctxt *cli.Context) error {
	infile := ""
	if ctxt.NArg() > 0 {
		infile = ctxt.Args()[0]
	}

	tree, _, err := helperRead(infile, "genea", ctxt.String("subset"))
	if err != nil {
		return err
	}

	result, err := tree.Query(ctxt)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", result)
	return nil
}
