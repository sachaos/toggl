package main

import (
	"fmt"
	"os"

	"github.com/sachaos/toggl/command"
	"github.com/urfave/cli"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "start",
		Usage:  "",
		Action: command.CmdStart,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "stop",
		Usage:  "",
		Action: command.CmdStop,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "current",
		Usage:  "",
		Action: command.CmdCurrent,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
