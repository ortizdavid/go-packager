package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)


type PackagerConfig struct {
	BuildName     string   `json:"build_name"`
	BuildTarget   string   `json:"build_target"`
	System        string   `json:"system"`
	PackagedItems []string `json:"packaged_items"`
}


func DecodeConfig(fileName string) PackagerConfig {
	var config PackagerConfig
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error while openning file:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatal("Error while decoding JSON", err)
	}

	return config
}


func (pkgConfig PackagerConfig) PrintConfigs() {
	fmt.Println("Build Name: ", pkgConfig.BuildName)
	fmt.Println("System: ", pkgConfig.System)
	fmt.Println("Build Target: ", pkgConfig.BuildTarget)
	fmt.Println("Packaged Items: ")
	for _, item := range pkgConfig.PackagedItems {
		fmt.Printf(" - %s\n", item)
	}
	fmt.Printf("\n")
}


func (pkgConfig PackagerConfig) CheckConfig(configFile string) {
	config := DecodeConfig(configFile)
	fmt.Println("Showing 'go-packager.json' configurations:")
	config.PrintConfigs()
}


