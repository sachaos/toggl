package command

import (
	"errors"
	"strconv"

	"github.com/sachaos/toggl/utils"
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

	utils.CreateConfig(utils.LocalConfigFilePath(), map[string]int{"wid": wid})

	return nil
}
