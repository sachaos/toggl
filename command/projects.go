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

func getProjects(c *cli.Context) (projects toggl.Projects, err error) {
	if c.Bool("cache")
	projects = cache.GetContent().Projects
	if len(projects) == 0 {
		projects, err = toggl.FetchWorkspaceProjects(viper.GetString("token"), viper.GetInt("wid"))
		if err != nil {
			return
		}
		cache.SetProjects(projects)
		cache.Write()
	}
	return
}

func CmdProjects(c *cli.Context) (err error) {
	var projects toggl.Projects
	projects, err = getProjects(c)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 4, 1, ' ', 0)

	for _, project := range projects {
		fmt.Fprintf(w, "%d\t%s\n",
			project.ID,
			project.Name,
		)
	}
	w.Flush()

	return
}
