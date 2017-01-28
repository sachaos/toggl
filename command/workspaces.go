package command

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/sachaos/toggl/cache"
	"github.com/sachaos/toggl/lib"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func GetWorkspaces(c *cli.Context) (workspaces toggl.Workspaces, err error) {
	workspaces = cache.GetContent().Workspaces
	if len(workspaces) == 0 {
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

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 4, 1, ' ', 0)

	for _, workspace := range workspaces {
		var flg string
		if workspace.ID == viper.GetInt("wid") {
			flg = "*"
		} else {
			flg = ""
		}
		fmt.Fprintf(w, "%s\t%d\t%s\n",
			flg,
			workspace.ID,
			workspace.Name,
		)
	}
	w.Flush()

	return nil
}
