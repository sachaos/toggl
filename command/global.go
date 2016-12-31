package command

import (
	"errors"
	"strconv"

	"github.com/sachaos/toggl/utils"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func CmdGlobal(c *cli.Context) error {
	if !c.Args().Present() {
		return errors.New("Command Failed")
	}
	wid, err := strconv.Atoi(c.Args().First())
	if err != nil {
		return err
	}
	viper.Set("wid", wid)

	utils.CreateConfig(utils.RootConfigFilePath(), viper.AllSettings())

	return nil
}
