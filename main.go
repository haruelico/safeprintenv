package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/pelletier/go-toml"
)

type Configure struct {
	SensitiveList SensitiveList
}

type SensitiveList struct {
	Keys []string
}

func main() {
	envs := os.Environ()

	conf, err := loadConfigure()
	if err != nil {
		fmt.Println(err)
	}

	printFilteredEnvVars(envs, conf)
}

func isSensitiveEnvVar(env string, conf *Configure) bool {
	givenKey := strings.SplitN(env, "=", 2)[0]
	for _, key := range conf.SensitiveList.Keys {
		if givenKey == key {
			return true
		}
	}
	return false
}

func loadConfigure() (*Configure, error) {
	conf := new(Configure)
	// confDir, err := os.UserConfigDir()

	// if err != nil {
	// 	return nil, err
	// }

	// confPath := path.Join(confDir, "safeprintenv", "config.toml")
	confPath := path.Join("config.example.toml")
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		// config file not exist. Return empty config
		return conf, nil
	}

	// rawConfBody, err := ioutil.ReadFile(path.Join(confPath))
	rawConfBody, _ := ioutil.ReadFile("config.example.toml")
	toml.Unmarshal(rawConfBody, conf)

	return conf, nil
}

func printFilteredEnvVars(envs []string, conf *Configure) {
	for _, env := range envs {
		if !isSensitiveEnvVar(env, conf) {
			fmt.Println(env)
		}
	}
}
