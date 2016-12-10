package command

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/sachaos/toggl/lib"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func CmdCurrent(c *cli.Context) error {

	current, err := toggl.FetchCurrent(viper.GetString("token"))
	current_time_entry := current.Data
	if err != nil {
		return err
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 4, 1, ' ', 0)

	fmt.Fprintf(w, "ID\t%d\n", current_time_entry.ID)
	fmt.Fprintf(w, "Description\t%s\n", current_time_entry.Description)
	w.Flush()

	return nil
}
