package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sachaos/toggl/cache"
	"github.com/sachaos/toggl/lib"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

var (
	ConfigPath = os.Getenv("HOME")
)

const (
	ConfigName = ".toggl"
	ConfigType = "json"
)

func main() {
	cache.New(os.Getenv("HOME") + "/.toggl.cache.json")
	cache.Init()

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

func RootConfigFilePath() string {
	return filepath.Join(ConfigPath, ConfigName+"."+ConfigType)
}

func LocalConfigFilePath() string {
	return filepath.Join(".", ConfigName+"."+ConfigType)
}

func CreateConfig(path string, hash interface{}) {
	buf, _ := json.MarshalIndent(hash, "", "  ")
	err := ioutil.WriteFile(path, buf, os.ModePerm)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func LoadLocalConfig() error {
	localFilePath := LocalConfigFilePath()
	file, err := os.Open(localFilePath)
	if err != nil {
		return err
	}
	viper.MergeConfig(file)
	return nil
}

func initialize() {
	viper.SetConfigType(ConfigType)
	viper.SetConfigName(ConfigName)
	viper.AddConfigPath(ConfigPath)
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	LoadLocalConfig()

	if !viper.IsSet("token") {
		requireToken()
		CreateConfig(RootConfigFilePath(), viper.AllSettings())
	}
}
