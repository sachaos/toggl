package command

import (
	"encoding/csv"
	"os"

	"github.com/sachaos/toggl/util"
	"github.com/urfave/cli"
)

func NewWriter(c *cli.Context) (writer util.Writer) {
	if c.GlobalBool("csv") {
		writer = csv.NewWriter(os.Stdout)
	} else {
		writer = util.NewTabWriter(os.Stdout)
	}
	return
}
