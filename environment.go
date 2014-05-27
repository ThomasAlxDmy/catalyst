package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

var (
	currentFolder string
)

// Tries to load a custom configuration file
// If it fails or the data is empty then loads the default one
func ConfigurationData() (data []byte) {
	if currentFolder != "" {
		data = CustomConfigFileData(currentFolder)
	}

	if data == nil {
		data = DefaultConfigFileData()
	}

	return data
}

// Tries to load a custom configuration file
func CustomConfigFileData(baseFolder string) []byte {
	filePath := path.Join(baseFolder, "catalyst_config.yml")
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("Can't load a custom config file `", filePath, "`. Error:", err)
	}

	return data
}

// Tries to load the default configuration file
func DefaultConfigFileData() []byte {
	_, filename, _, _ := runtime.Caller(1)
	filePath := path.Join(path.Dir(filename), "catalyst_config.yml")

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error reading default file. Catalyst can't work without a config file. Error:", err)
	}

	return data
}

func GetTreeDirectory() []string {
	currentFolder = getWatchingFolder()

	return buildTreeDirectory(currentFolder)
}

func getWatchingFolder() string {
	folder, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return folder
}

func getPackageFolder() string {
	src_gopath := os.Getenv("GOPATH") + "/src/"
	if !strings.HasPrefix(currentFolder, src_gopath) {
		log.Println("Warning the current repository is not part of the GOPATH:", src_gopath)
	}

	return strings.TrimPrefix(currentFolder, src_gopath)
}

// Recursively builds the folder tree directory based of a first path in parameter
func buildTreeDirectory(baseFolderPath string) []string {
	files, err := ioutil.ReadDir(baseFolderPath)
	if err != nil {
		log.Fatal("Error while accessing `", baseFolderPath, "`. ", err)
	}

	filesPath := []string{baseFolderPath}
	for _, file := range files {
		if file.IsDir() && !strings.HasPrefix(file.Name(), ".") {
			folderPath := baseFolderPath + "/" + file.Name()
			filesPath = append(filesPath, buildTreeDirectory(folderPath)...)
		}
	}

	return filesPath
}
