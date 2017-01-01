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

	CreateConfig(LocalConfigFilePath(), map[string]int{"wid": wid})

	return nil
}
