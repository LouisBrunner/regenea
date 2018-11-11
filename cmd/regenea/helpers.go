package main

import (
	"fmt"
	"os"

	"github.com/LouisBrunner/regenea/genea"
)

const (
	formGenea = "genea"
)

func helperRead(infile, format string) (*genea.Tree, genea.Version, error) {
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

	var tree *genea.Tree
	var version genea.Version
	switch format {
	case formGenea:
		tree, version, err = genea.Parse(in)
		if err != nil {
			return nil, 0, err
		}
	default:
		return nil, 0, fmt.Errorf("unknown format: `%s`", format)
	}

	return tree, version, err
}
