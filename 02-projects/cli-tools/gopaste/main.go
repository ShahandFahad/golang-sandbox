/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/ShahandFahad/golang-sandbox/02-projects/cli-tools/gopaste/cmd"
	"github.com/ShahandFahad/golang-sandbox/02-projects/cli-tools/gopaste/config"
)

func main() {
	config.LoadConfig()
	cmd.Execute()
}
