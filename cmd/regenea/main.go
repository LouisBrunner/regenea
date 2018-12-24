package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

// TODO: add display/search for person

func getCLI() *cli.App {
	app := cli.NewApp()
	app.Name = "regenea"
	app.EnableBashCompletion = true
	app.HideVersion = true

	app.Commands = []cli.Command{
		getCheckCommand(),
		getTransformCommand(),
		getQueryCommand(),
		getReportCommand(),
		getDisplayCommand(),
	}

	return app
}

func main() {
	app := getCLI()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
