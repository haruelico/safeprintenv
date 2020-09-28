package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/pelletier/go-toml"
	flag "github.com/spf13/pflag"
)

type Configure struct {
	SensitiveList SensitiveList
	Filterstyle   Filterstyle
}

type SensitiveList struct {
	Keys []string
}

type Filterstyle struct {
	Style string
}

func main() {
	envs := os.Environ()
	var showAll *bool = flag.Bool("unsafe-all", false, "Show all environment variables. No filtered.")
	flag.Parse()

	conf, err := loadConfigure()
	if err != nil {
		fmt.Println(err)
	}

	if *showAll {
		for _, env := range envs {
			fmt.Println(env)
		}
	} else {
		printFilteredEnvVars(envs, conf)
	}
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
	confDir, err := os.UserConfigDir()

	if err != nil {
		return nil, err
	}

	confPath := path.Join(confDir, "safeprintenv", "config.toml")
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		// config file not exist. Return empty config
		return conf, nil
	}

	rawConfBody, err := ioutil.ReadFile(path.Join(confPath))
	toml.Unmarshal(rawConfBody, conf)

	return conf, nil
}

func printFilteredEnvVars(envs []string, conf *Configure) {
	for _, env := range envs {
		if !isSensitiveEnvVar(env, conf) {
			fmt.Println(env)
		} else if conf.Filterstyle.Style == "masked" {
			givenKey := strings.SplitN(env, "=", 2)[0]
			fmt.Printf("%s=********\n", givenKey)
		}
	}
}
