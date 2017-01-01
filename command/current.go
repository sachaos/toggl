package command

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/sachaos/toggl/lib"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func formatTimeDuration(duration time.Duration) string {
	hours := duration / time.Hour
	minutes := duration / time.Minute % 60
	seconds := duration / time.Second % 60
	return fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds)
}

func calcDuration(duration int64) time.Duration {
	second := duration + time.Now().Unix()
	return time.Duration(second * int64(time.Second))
}

func CmdCurrent(c *cli.Context) error {
	var project toggl.Project

	current, err := toggl.GetCurrentTimeEntry(viper.GetString("token"))
	current_time_entry := current.Data
	if err != nil {
		return err
	}
	workspaces, err := toggl.FetchWorkspaces(viper.GetString("token"))
	if err != nil {
		return err
	}
	workspace, err := workspaces.FindByID(current_time_entry.WID)

	if current_time_entry.ID == 0 {
		fmt.Println("No time entry")
		return nil
	}

	if current_time_entry.PID != 0 {
		projects, err := toggl.FetchWorkspaceProjects(viper.GetString("token"), current_time_entry.WID)
		if err != nil {
			return err
		}
		project, err = projects.FindByID(current_time_entry.PID)
		if err != nil {
			return err
		}
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 4, 1, ' ', 0)

	fmt.Fprintf(w, "ID\t%d\n", current_time_entry.ID)
	fmt.Fprintf(w, "Description\t%s\n", current_time_entry.Description)
	fmt.Fprintf(w, "Project\t%s\n", project.Name)
	fmt.Fprintf(w, "Workspace\t%s\n", workspace.Name)
	fmt.Fprintf(w, "Duration\t%s\n", formatTimeDuration(calcDuration(current_time_entry.Duration)))
	w.Flush()

	return nil
}
