package main

import (
	"fmt"
	"os"

	"github.com/sachaos/toggl/lib"
	"github.com/sachaos/toggl/utils"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

func main() {
	initialize()

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "sachaos"
	app.Email = "sakataku7@gmail.com"
	app.Usage = "Toggl API CLI Client"

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}

func requireToken() error {
	var token string
	var workspaces []toggl.Workspace
	var err error
	count := 0
	for count < 3 {
		fmt.Printf("Input API Token: ")
		fmt.Scan(&token)
		workspaces, err = toggl.FetchWorkspaces(token)
		if err == nil {
			viper.Set("token", token)
			viper.Set("wid", workspaces[0].ID)
			return nil
		}
		count++
	}
	panic(fmt.Errorf("Invalid token"))
}

func LoadLocalConfig() error {
	localFilePath := utils.LocalConfigFilePath()
	file, err := os.Open(localFilePath)
	if err != nil {
		return err
	}
	viper.MergeConfig(file)
	return nil
}

func initialize() {
	viper.SetConfigType(utils.ConfigType)
	viper.SetConfigName(utils.ConfigName)
	viper.AddConfigPath(utils.ConfigPath)
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	LoadLocalConfig()

	if !viper.IsSet("token") {
		requireToken()
		utils.CreateConfig(utils.RootConfigFilePath(), viper.AllSettings())
	}
}
