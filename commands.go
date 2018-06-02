package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var GlobalFlags = []cli.Flag{
	cli.BoolFlag{
		Name: "cache",
	},
	cli.BoolFlag{
		Name: "csv",
	},
}

var projectIDFlag = cli.IntFlag{
	Name:  "project-id, P",
	Usage: "Project id",
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
