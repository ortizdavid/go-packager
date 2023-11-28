package builders

import (
	"fmt"
	"log"
	"os/exec"
	"github.com/ortizdavid/go-nopain/filemanager"
	"github.com/ortizdavid/go-packager/config"
)

type AppBuilder struct {
}

var zipper filemanager.FileZip
var fileManager filemanager.FileManager
var fileInfo filemanager.FileInfo


func (builder *AppBuilder) Run(configFile string) {
	pkgConfig := config.DecodeConfig(configFile)
	log.Printf("\nPackaging the application '%s'", pkgConfig.BuildName)
	builder.copyItems(pkgConfig)
	builder.buildProject(pkgConfig)
	builder.zipFolder(pkgConfig)
	builder.removeOldFolder(pkgConfig)
}


func (builder *AppBuilder) copyItems(pkgConfig config.PackagerConfig) {
	dest := pkgConfig.BuildTarget +"/"+ pkgConfig.BuildName +"/"
	for _, src := range pkgConfig.PackagedItems {
		if fileInfo.IsDir(src) {
			fileManager.CopyFolder(src, dest +"/"+ src)
			fmt.Printf("\nCopying Folder '%s'", src)
		} else {
			fileManager.CopyFile(src, dest +"/"+ src)
			fmt.Printf("\nCopying File '%s'", src)
		}
	}
}


func (builder *AppBuilder) buildProject(pkgConfig config.PackagerConfig) {
	strBuild := "build"
	outputPath := "-o"
	outputFile := pkgConfig.BuildTarget +"/"+ pkgConfig.BuildName +"/"+ builder.getOutPutFile(pkgConfig)
	
	log.Printf("\nBuilding App '%s'", pkgConfig.BuildName)
	cmd := exec.Command("go", strBuild, outputPath, outputFile)
	output, err := cmd.Output()

	log.Println(output)
	if err != nil {
		log.Fatal("Error:", err)
	}
	log.Println("Builded Sucessuly!")
}


func (builder *AppBuilder) zipFolder(pkgConfig config.PackagerConfig) {
	source := pkgConfig.BuildTarget +"/"+ pkgConfig.BuildName
	target := pkgConfig.BuildTarget +"/"+ pkgConfig.BuildName +".zip"
	log.Printf("\nCompressing Folder '%s'", target)
	err := zipper.ZipFileOrFolder(source, target)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("\nCompressed successfuly!")
}

func (builder *AppBuilder) removeOldFolder(pkgConfig config.PackagerConfig) {
	folder := pkgConfig.BuildTarget +"/"+ pkgConfig.BuildName
	
	log.Printf("\nRemoving Old Folder '%s'", folder)
	removed := fileManager.RemoveFolder(folder)
	if !removed {
		log.Fatal("Error removing folder")
	}
	log.Println("\nRemoved successfuly!")
}


func (builder *AppBuilder) getOutPutFile(pkgConfig config.PackagerConfig) string {
	buildName := pkgConfig.BuildName
	outputFile := buildName
	if pkgConfig.System == "windows" {
		outputFile = buildName + ".exe"
	}
	return outputFile
}


