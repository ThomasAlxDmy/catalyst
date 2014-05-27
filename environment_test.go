package main

import (
	"io/ioutil"
	"log"
	"path"
	"reflect"
	"testing"
)

var expectedYamlData []byte

func init() {
	var err error

	expectedYamlData, err = ioutil.ReadFile(path.Join(currentFolder, "catalyst_config.yml"))
	if err != nil {
		log.Fatalf("Error The yaml file was not found", err)
	}
}

func TestGetWatchingFolder(t *testing.T) {
	packagePath := getWatchingFolder()

	if packagePath == "" {
		t.Fatal("Error when trying to find to full path to the package.")
	}
}

func TestGetPackageFolder(t *testing.T) {
	packagePath := getPackageFolder()
	if packagePath != "github.com/ThomasAlxDmy/catalyst" {
		t.Fatal("Error the package path has an unexpected value: ", packagePath)
	}
}

func TestGetTreeDirectory(t *testing.T) {
	baseFolder := getWatchingFolder()
	folders := []string{"", "/command"}
	folderTree := GetTreeDirectory()

	if currentFolder == "" {
		t.Fatal("currentFolder has not been initialzed")
	}

	for index, folder := range folders {
		folders[index] = baseFolder + folder
	}

	if !reflect.DeepEqual(folderTree, folders) {
		t.Fatal("The directory structure is not correct. got:", folderTree, "expected:", folders)
	}
}

func TestConfigurationData(t *testing.T) {
	yamlData := ConfigurationData()

	if !reflect.DeepEqual(yamlData, expectedYamlData) {
		t.Fatalf("Error yaml file does not have the expected value got: ", yamlData, ". Expected:", expectedYamlData)
	}
}

func TestCustomConfigFileData(t *testing.T) {
	yamlData := CustomConfigFileData("doesNotExist")
	if yamlData != nil {
		t.Fatalf("Error when loading a file that does not exist. value expected nil got:", yamlData)
	}

	yamlData = CustomConfigFileData(currentFolder)
	if !reflect.DeepEqual(yamlData, expectedYamlData) {
		t.Fatalf("Error yaml file does not have the expected value got: ", yamlData, ". Expected:", expectedYamlData)
	}
}

func TestDefaultConfigFileData(t *testing.T) {
	yamlData := DefaultConfigFileData()

	if !reflect.DeepEqual(yamlData, expectedYamlData) {
		t.Fatalf("Error yaml file does not have the expected value got: ", yamlData, ". Expected:", expectedYamlData)
	}
}
