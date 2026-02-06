package main

import (
	"Aj-vrod/bicho/cmd"
)

var version = "dev-0.0.0"

func main() {
	rootCmd := cmd.NewRootCmd(version)
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
