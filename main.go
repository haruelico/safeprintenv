package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	envs := os.Environ()

	for _, env := range envs {
		fmt.Println(isSnsitiveEnvVar(env))
	}
}

func isSnsitiveEnvVar(env string) bool {
	key := strings.SplitN(env, "=", 2)[0]
	if key == "SAMPLE_SECRET_ENV" {
		return true
	}
	return false
}
