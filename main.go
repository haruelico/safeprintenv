package main

import (
	"fmt"
	"os"
)

func main() {
	envs := os.Environ()

	for _, env := range envs {
		fmt.Println(env)
	}
}
