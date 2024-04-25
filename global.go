package main

import (
	"errors"
	"strconv"

	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func CmdGlobal(c *cli.Context) error {
	if !c.Args().Present() {
		return errors.New("Command Failed")
	}
	workspaceID, err := strconv.Atoi(c.Args().First())
	if err != nil {
		return err
	}
	viper.Set("wid", workspaceID)

	CreateConfig(RootConfigFilePath(), viper.AllSettings())

	return nil
}
