package main

import (
	"log"
	"os"

	cli "gopkg.in/urfave/cli.v1"

	// Commands.
	"github.com/itsfunky/funky/cmd/funky/serve"

	// Providers.
	_ "github.com/itsfunky/funky/providers/aws"
)

var (
	app = cli.NewApp()

	// Version will be set using ldflags on CI
	Version = "development"
)

func init() {
	app.Name = "funky"
	app.HelpName = "funky"
	app.Usage = "Create and manage cross-cloud functions and resources."
	app.UsageText = ""
	app.EnableBashCompletion = true
	app.Version = Version

	app.Commands = []cli.Command{
		{
			Name:    "serve",
			Aliases: []string{"s"},
			Usage:   "serve your functions on a local server",
			Action: func(_ *cli.Context) {
				serve.Serve()
			},
		},
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
