package main

import (
	"fmt"
	"os"

	"github.com/LouisBrunner/regenea/genea"

	"github.com/Songmu/prompter"
	"gopkg.in/urfave/cli.v1"
)

const defaultOutVersion = 0

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
				Usage: "output file where to write the transformed file (default to stdout)",
			},
			cli.StringFlag{
				Name:  "outform",
				Usage: "format of the input file",
			},
			cli.UintFlag{
				Name:  "outversion",
				Usage: "when using `outform` with `genea`, you can specify which version should be output",
				Value: defaultOutVersion,
			},
			cli.BoolFlag{
				Name:  "pretty",
				Usage: "prettify the output",
			},
		},
		Action: doTransformCommand,
	}
}

func confirmInformationLoss(conversion string) bool {
	msg := fmt.Sprintf(
		"There is a loss of information going %s, are you sure you want to convert?",
		conversion,
	)
	return prompter.YN(msg, false)
}

func doTransformCommand(ctxt *cli.Context) error {
	format := ctxt.String("outform")
	if format == "" {
		return fmt.Errorf("missing output format")
	}

	inform := ctxt.String("inform")
	tree, inversion, err := helperRead(ctxt.String("in"), inform)
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

	switch format {
	case formGenea:
		version := genea.VersionV2
		switch versionRaw := ctxt.Uint("outversion"); versionRaw {
		case defaultOutVersion:
		case 1:
			if inform == formGenea && inversion == genea.VersionV2 {
				if !confirmInformationLoss("from genea V2 to genea V1") {
					return fmt.Errorf("transformation aborted")
				}
			}
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
