package command

import (
	"github.com/sachaos/toggl/cache"
	"github.com/sachaos/toggl/lib"
	"github.com/urfave/cli"
)

func (app *App) CmdStop(c *cli.Context) error {
	current, err := app.client.GetCurrentTimeEntry()
	current_time_entry := current.Data

	err = app.client.PutStopTimeEntry(current_time_entry.ID)

	if err != nil {
		return err
	}

	cache.SetCurrentTimeEntry(toggl.TimeEntry{})
	cache.Write()

	return nil
}
