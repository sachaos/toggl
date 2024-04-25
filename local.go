package main

import (
	"errors"
	"strconv"

	"github.com/urfave/cli"
)

func CmdLocal(c *cli.Context) error {
	if !c.Args().Present() {
		return errors.New("Command Failed")
	}
	workspaceID, err := strconv.Atoi(c.Args().First())
	if err != nil {
		return err
	}

	var projectID int

	if c.IsSet("project-id") {
		projectID = c.Int("project-id")
	}

	CreateConfig(LocalConfigFilePath(), map[string]int{"wid": workspaceID, "pid": projectID})

	return nil
}
