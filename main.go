package main

import (
	"github.com/Israel-Ferreira/stock-challenge-cli/cmd"
	"github.com/Israel-Ferreira/stock-challenge-cli/config"
)

func main() {
	config.LoadEnv()

	cmd.RootCmd()
}
