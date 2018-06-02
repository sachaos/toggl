package command

import (
	"errors"

	"github.com/sachaos/toggl/cache"
	"github.com/sachaos/toggl/lib"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func (app *App) CmdStart(c *cli.Context) error {
	timeEntry := toggl.TimeEntry{}
	if !c.Args().Present() {
		return errors.New("Command Failed")
	}

	timeEntry.Description = c.Args().First()
	timeEntry.WID = viper.GetInt("wid")
	if c.IsSet("project-id") {
		timeEntry.PID = c.Int("project-id")
	}
	response, err := app.client.PostStartTimeEntry(timeEntry)
	if err != nil {
		return err
	}

	cache.SetCurrentTimeEntry(response.Data)
	cache.Write()

	return nil
}
