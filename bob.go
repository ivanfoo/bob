package main

import (
	//"fmt"
	"os"

	"github.com/ivanfoo/bob/cmd"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:   "get",
			Usage:  "download repository",
			Action: cmd.Get,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "u, update",
					Usage: "refresh local copy",
				},
			},
		},
		{
			Name:   "init",
			Usage:  "set up remotes",
			Action: cmd.Init,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "t, type",
					Usage: "kind of tech",
				},
			},
		},
		{
			Name:   "publish",
			Usage:  "publish changes to forge",
			Action: cmd.Publish,
		},
		{
			Name:   "sync",
			Usage:  "sync remotes with local changes",
			Action: cmd.Sync,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "m, mine",
					Usage: "remote branch of the project",
				},
				cli.StringFlag{
					Name:  "f, forge",
					Usage: "remote branch of the forge",
				},
				cli.BoolFlag{
					Name:  "pr, pull-request",
					Usage: "make a pr to forge's master branch",
				},
			},
		},
	}

	app.Run(os.Args)
}
