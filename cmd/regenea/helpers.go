package main

import (
	"fmt"
	"os"

	"github.com/LouisBrunner/regenea/genea"

	"gopkg.in/urfave/cli.v1"
)

const (
	formGenea = "genea"
)

var inField = cli.StringFlag{
	Name:  "in",
	Usage: "input file to transform (default to stdin)",
}

var subsetField = cli.StringFlag{
	Name:  "subset",
	Usage: "use a subset of the tree including the given person, their ascendants, descendants and close family (partners, siblings, cousins, uncles and aunts)",
}

func helperRead(infile, format, subset string) (*genea.Tree, genea.Version, error) {
	var err error
	in := os.Stdin

	if format == "" {
		return nil, 0, fmt.Errorf("missing input format")
	}

	if infile != "" {
		in, err = os.Open(infile)
		if err != nil {
			return nil, 0, err
		}
	}

	switch format {
	case formGenea:
		tree, version, err := genea.Parse(in)
		if err == nil && subset != "" {
			tree, err = tree.Subtree(subset)
		}
		return tree, version, err
	default:
		return nil, 0, fmt.Errorf("unknown format: `%s`", format)
	}
}
