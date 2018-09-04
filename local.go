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
	wid, err := strconv.Atoi(c.Args().First())
	if err != nil {
		return err
	}

	var pid int

	if c.IsSet("project-id") {
		pid = c.Int("project-id")
	}

	CreateConfig(LocalConfigFilePath(), map[string]int{"wid": wid, "pid": pid})

	return nil
}
