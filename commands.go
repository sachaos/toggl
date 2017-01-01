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
		Usage:  "Start time entry",
		Action: command.CmdStart,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "stop",
		Usage:  "End time entry",
		Action: command.CmdStop,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "current",
		Usage:  "Show current time entry",
		Action: command.CmdCurrent,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "workspaces",
		Usage:  "Show workspaces",
		Action: command.CmdWorkspaces,
	},
	{
		Name:   "projects",
		Usage:  "Show projects on current workspaces",
		Action: command.CmdProjects,
	},
	{
		Name:   "local",
		Usage:  "Set current dir workspace",
		Action: CmdLocal,
	},
	{
		Name:   "global",
		Usage:  "Set global workspace",
		Action: CmdGlobal,
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
