package main

import (
	"os"
	"github.com/soypita/clirescue/trackerapi"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "clirescue"
	app.Usage = "CLI tool to talk to the Pivotal Tracker's API"

	app.Commands = []cli.Command{
		{
			Name:  "me",
			Usage: "prints out Tracker's representation of your account",
			Action: func(c *cli.Context) {
				trackerapi.Me()
			},
		},
		{
			Name:  "projects",
			Usage: "prints out Tracker's projects",
			Action: func(c *cli.Context) {
				trackerapi.Projects()
			},
		},
	}

	app.Run(os.Args)
}
