package command

import (
	"errors"

	"github.com/sachaos/toggl/cache"
	toggl "github.com/sachaos/toggl/lib"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func (app *App) CmdStart(c *cli.Context) error {
	timeEntry := toggl.TimeEntry{}
	if !c.Args().Present() {
		return errors.New("Command Failed")
	}

	timeEntry.Description = c.Args().First()
	timeEntry.WorkspaceID = viper.GetInt("wid")
	if c.IsSet("project-id") {
		timeEntry.ProjectID = c.Int("project-id")
	} else if viper.GetInt("pid") != 0 {
		timeEntry.ProjectID = viper.GetInt("pid")
	}

	newTimeEntry, err := app.client.PostStartTimeEntry(timeEntry)
	if err != nil {
		return err
	}

	cache.SetCurrentTimeEntry(newTimeEntry)
	cache.Write()

	return nil
}
