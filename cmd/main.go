package main

import (
	"os"

	"github.com/ortizdavid/go-packager/helpers"
	"github.com/ortizdavid/go-packager/builders"
)


func main() {
	cliArgs := os.Args
	numArgs := len(cliArgs)
	
	if numArgs > 1 {
		if numArgs == 2 && cliArgs[1] == "-help" {
			helpers.PrintHelp()
		} else if numArgs == 2 && cliArgs[1] == "-examples" {
			helpers.PrintExamples()
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