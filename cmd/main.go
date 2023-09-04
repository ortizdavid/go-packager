package main

import (
	"os"
	"github.com/ortizdavid/go-packager/builders"
	"github.com/ortizdavid/go-packager/config"
	"github.com/ortizdavid/go-packager/helpers"
)


func main() {
	cliArgs := os.Args
	numArgs := len(cliArgs)
	
	if numArgs > 1 {
		if numArgs == 2 && cliArgs[1] == "-help" {
			helpers.PrintHelp()
		} else if numArgs == 2 && cliArgs[1] == "-examples" {
			helpers.PrintExamples()
		} else if numArgs == 2 && cliArgs[1] == "-check" {
			config.PackagerConfig{}.CheckConfig("go-packager.json")
		} else if numArgs == 2 && cliArgs[1] == "run" {
			var builder builders.AppBuilder
			builder.Run("go-packager.json")
		} else {
			helpers.PrintHelp()
		}
	} else {
		helpers.PrintHelp()
	}
}