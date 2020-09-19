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

	for _, env := range envs {
		fmt.Println(isSnsitiveEnvVar(env))
	}
	conf, err := loadConfigure()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conf)
}

func isSnsitiveEnvVar(env string) bool {
	key := strings.SplitN(env, "=", 2)[0]
	if key == "SAMPLE_SECRET_ENV" {
		return true
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
