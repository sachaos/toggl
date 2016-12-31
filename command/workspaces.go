package command

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/sachaos/toggl/lib"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func CmdWorkspaces(c *cli.Context) error {
	workspaces, err := toggl.FetchWorkspaces(viper.GetString("token"))
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
