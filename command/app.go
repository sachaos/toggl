package command

import (
	"github.com/sachaos/toggl/lib"
)

type App struct {
	client *toggl.Client
}

func NewApp(token string) *App {
	return &App{
		client: toggl.NewDefaultClient(token),
	}
}
