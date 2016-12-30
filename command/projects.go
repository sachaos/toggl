package command

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/sachaos/toggl/lib"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func CmdProjects(c *cli.Context) error {
	projects, err := toggl.FetchWorkspaceProjects(viper.GetString("token"), viper.GetInt("wid"))
	if err != nil {
		return err
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 4, 1, ' ', 0)

	for _, project := range projects {
		fmt.Fprintf(w, "%d\t%s\n",
			project.ID,
			project.Name,
		)
	}
	w.Flush()

	return nil
}
