package command

import (
	"errors"

	"github.com/sachaos/toggl/lib"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func CmdStart(c *cli.Context) error {
	timeEntry := toggl.TimeEntry{}
	if !c.Args().Present() {
		return errors.New("Command Failed")
	}

	timeEntry.Description = c.Args().First()
	err := toggl.PostStartTimeEntry(timeEntry, viper.GetString("token"))
	if err != nil {
		return err
	}

	return nil
}
