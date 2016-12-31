package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	ConfigPath = os.Getenv("HOME")
)

const (
	ConfigName = ".toggl"
	ConfigType = "json"
)

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
