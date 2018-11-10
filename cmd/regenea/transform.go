package main

import (
	"fmt"
	"os"

	"github.com/LouisBrunner/regenea/genea"

	"gopkg.in/urfave/cli.v1"
)

func getTransformCommand() cli.Command {
	return cli.Command{
		Name:  "transform",
		Usage: "perform a transformation between 2 formats or 2 versions of the same format (supported: genea)",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "in",
				Usage: "input file to transform (default to stdin)",
			},
			cli.StringFlag{
				Name:  "inform",
				Usage: "format of the input file",
			},
			cli.StringFlag{
				Name:  "out",
				Usage: "output file where to write the transformed file  (default to stdout)",
			},
			cli.StringFlag{
				Name:  "outform",
				Usage: "format of the input file",
			},
			cli.UintFlag{
				Name:  "outversion",
				Usage: "when using `outform` with `genea`, you can specify which version should be output",
			},
			cli.BoolFlag{
				Name:  "pretty",
				Usage: "prettify the output",
			},
		},
		Action: doTransformCommand,
	}
}

func doTransformCommand(ctxt *cli.Context) error {
	tree, err := helperRead(ctxt.String("in"), ctxt.String("inform"))
	if err != nil {
		return err
	}

	out := os.Stdout
	if outfile := ctxt.String("out"); outfile != "" {
		out, err = os.Create(outfile)
		if err != nil {
			return err
		}
	}

	switch format := ctxt.String("outform"); format {
	case "genea":
		version := genea.VersionV2
		switch versionRaw := ctxt.Uint("outversion"); versionRaw {
		case 1:
			version = genea.VersionV1
		case 2:
		default:
			return fmt.Errorf("unknown version: `%d`", versionRaw)
		}
		err = tree.Format(out, version)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown format: `%s`", format)
	}

	return nil
}
