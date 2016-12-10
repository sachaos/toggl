package main

import (
	"fmt"
	"os"

	"encoding/json"
	"io/ioutil"
	"path/filepath"

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

	viper.SetConfigType(configType)
	viper.SetConfigName(configName) // name of config file (without extension)
	viper.AddConfigPath(configPath) // call multiple times to add many search paths
	viper.AddConfigPath(".")        // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file

	if err != nil {
		var token string
		fmt.Printf("Input API Token: ")
		fmt.Scan(&token)
		viper.Set("token", token)
		buf, err := json.MarshalIndent(viper.AllSettings(), "", "  ")
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		err = ioutil.WriteFile(filepath.Join(configPath, configName+"."+configType), buf, os.ModePerm)
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}

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
