package command

import (
	"github.com/sachaos/toggl/cache"
	"github.com/sachaos/toggl/lib"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func CmdStop(c *cli.Context) error {
	current, err := toggl.GetCurrentTimeEntry(viper.GetString("token"))
	current_time_entry := current.Data

	err = toggl.PutStopTimeEntry(current_time_entry.ID, viper.GetString("token"))

	if err != nil {
		return err
	}

	cache.SetCurrentTimeEntry(toggl.TimeEntry{})
	cache.Write()

	return nil
}
