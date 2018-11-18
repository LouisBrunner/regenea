package main

import (
	"encoding/json"
	"os"

	"github.com/LouisBrunner/regenea/report"

	"gopkg.in/urfave/cli.v1"
)

func getReportCommand() cli.Command {
	return cli.Command{
		Name:  "report",
		Usage: "display different stats about your tree (if no file is specified, stdin is used)",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "in",
				Usage: "input file to transform (default to stdin)",
			},
		},
		Action: doReportCommand,
	}
}

func doReportCommand(ctxt *cli.Context) error {
	infile := ""
	if ctxt.NArg() > 0 {
		infile = ctxt.Args()[0]
	}

	tree, _, err := helperRead(infile, "genea")
	if err != nil {
		return err
	}

	return stats.Process(tree, os.Stdout, func(v interface{}) ([]byte, error) {
		return json.MarshalIndent(v, "", "  ")
	})
}
