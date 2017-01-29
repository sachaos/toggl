package command

import (
	"strconv"

	"github.com/sachaos/toggl/cache"
	"github.com/sachaos/toggl/lib"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func GetWorkspaces(c *cli.Context) (workspaces toggl.Workspaces, err error) {
	workspaces = cache.GetContent().Workspaces
	if len(workspaces) == 0 || !c.GlobalBool("cache") {
		workspaces, err = toggl.FetchWorkspaces(viper.GetString("token"))
		cache.SetWorkspaces(workspaces)
		cache.Write()
	}
	return
}

func CmdWorkspaces(c *cli.Context) error {
	workspaces, err := GetWorkspaces(c)
	if err != nil {
		return err
	}

	writer := NewWriter(c)

	defer writer.Flush()

	for _, workspace := range workspaces {
		var flg string
		if workspace.ID == viper.GetInt("wid") {
			flg = "*"
		} else {
			flg = ""
		}
		writer.Write([]string{flg, strconv.Itoa(workspace.ID), workspace.Name})
	}

	return nil
}
