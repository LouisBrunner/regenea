package main

import (
	"log"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func getCLI() *cli.App {
	app := cli.NewApp()
	app.Name = "regenea"
	app.EnableBashCompletion = true
	app.HideVersion = true

	app.Commands = []cli.Command{
		getCheckCommand(),
		getTransformCommand(),
		// getStatsCommand(),
		// getDisplayCommand(),
	}

	return app
}

func main() {
	app := getCLI()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
