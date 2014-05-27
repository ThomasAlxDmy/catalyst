package main

import (
	"os"
	"testing"

	"github.com/ThomasAlxDmy/catalyst/command"
)

var (
	CommandCreateTmpFile = command.Command{Name: "touch", Arguments: []string{"tmp"}}
	CommandDeleteTmpFile = command.Command{Name: "rm", Arguments: []string{"tmp"}}
)

func TestCheckProjectHealth(t *testing.T) {
	watchFolders(folderTree)
	checkProjectHealth([]command.Command{CommandCreateTmpFile})
	_, fileError := os.Stat("tmp")

	if fileError != nil {
		t.Fatal("Error while executing commands:", fileError)
	}

	checkProjectHealth([]command.Command{CommandDeleteTmpFile})
}
