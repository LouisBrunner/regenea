package main

import (
	"fmt"
	"os"

	"github.com/LouisBrunner/regenea/genea"
)

func helperRead(infile, format string) (*genea.Tree, error) {
	var err error
	in := os.Stdin

	if infile != "" {
		in, err = os.Open(infile)
		if err != nil {
			return nil, err
		}
	}

	var tree *genea.Tree
	switch format {
	case "genea":
		tree, err = genea.Parse(in)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown format: `%s`", format)
	}

	return tree, err
}
