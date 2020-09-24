package main

import (
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/urfave/cli/v2"

	"github.com/wesleimp/sfs/cmd/server"
)

var (
	version = "0.1.0"
)

func main() {
	log.SetHandler(text.Default)

	app := &cli.App{
		Name:     "sfs",
		HelpName: "sfs",
		Usage:    "A simple static file serving command-line tool.",
		Version:  version,
		Authors: []*cli.Author{
			{
				Name:  "Weslei Juan Novaes Pereira",
				Email: "wesleimsr@gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "Application port",
				Value:   "8800",
			},
		},
		Action: server.Start,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.WithError(err).Fatal("error running server")
	}

}
