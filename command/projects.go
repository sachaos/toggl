package command

import (
	"strconv"

	"github.com/sachaos/toggl/cache"
	"github.com/sachaos/toggl/lib"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func GetProjects(c *cli.Context) (projects toggl.Projects, err error) {
	projects = cache.GetContent().Projects
	if len(projects) == 0 || !c.GlobalBool("cache") {
		projects, err = toggl.FetchWorkspaceProjects(viper.GetString("token"), viper.GetInt("wid"))
		cache.SetProjects(projects)
		cache.Write()
	}
	return
}

func CmdProjects(c *cli.Context) error {
	projects, err := GetProjects(c)
	if err != nil {
		return err
	}

	writer := NewWriter(c)

	defer writer.Flush()

	for _, project := range projects {
		writer.Write([]string{strconv.Itoa(project.ID), project.Name})
	}

	return nil
}
