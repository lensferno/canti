package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	cli.AppHelpTemplate = helpTemplate

	app := &cli.App{
		Flags:    globalFlags,
		Commands: allCommands,
		Version:  "0.0.1-dev",
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
