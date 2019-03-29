package main

import (
	"log"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

var debug bool

func main() {
	app := cli.NewApp()
	app.Name = "sync-core"
	app.Usage = "A Mirror Job Management Tool"
	app.Version = "0.1.0"
	app.Author = "example"
	app.Email = "example@example.com"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "enable debug",
			Destination: &debug,
		},
	}
	app.Commands = []cli.Command{}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
