package main

import (
	"fmt"
	"os"

	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/sachaos/toggl/lib"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

var (
	configPath = os.Getenv("HOME")
)

const (
	configName = ".toggl"
	configType = "json"
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

func createConfig() {
	buf, _ := json.MarshalIndent(viper.AllSettings(), "", "  ")
	err := ioutil.WriteFile(filepath.Join(configPath, configName+"."+configType), buf, os.ModePerm)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func initialize() {
	viper.SetConfigType(configType)
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(".")
	_ = viper.ReadInConfig()

	if !viper.IsSet("token") {
		requireToken()
		createConfig()
	}
}
