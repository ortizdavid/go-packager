package helpers

import "fmt"

func PrintHelp() {
	help := `Tool to Build a Golang project, incluing all needed files to run the final application.

	USAGE:
		go-packager <COMMAND>

	COMMAND:
		-help: Show Help for use the tool
		-run: Runs go-packager, reading the 'packager.json' file
		-examples: Shows examples
	`
	fmt.Println(help)
}

func PrintExamples() {
	sample := `packager.json sample:
	{
		"build_name": "first-build",
		"build_target": "C:\\",  
		"system": "windows",
		"packaged_items": [
			"static",
			"templates",
			"logs",
			".env"
		]
	}
	`
	fmt.Print(sample)
	examples := `
	Command to run:
	go-packager run
	`
	fmt.Print(examples)
}