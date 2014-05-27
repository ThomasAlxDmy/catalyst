package main

import (
	"log"

	"github.com/ThomasAlxDmy/catalyst/command"

	"github.com/howeyc/fsnotify"
)

// folderTree is used to load the folder structure from where the application is launched
// watcher a pointer to a directory watcher
var (
	folderTree []string
	watcher    *fsnotify.Watcher
)

// FolderAction is a generic type of closure that matches:
// Take a string in parameters and without returns value
// It's used later in the code to dertermine what kind of action will be executed on a folder
type FolderAction func(string)

// Initializes the current tree directory and the watcher
func init() {
	var err error
	folderTree = GetTreeDirectory()

	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
}

// Watches all folder in the tree directory
// Loads the action to perform from the configuration
// Listens for actions to be performed on each of the file of the tree directory
func main() {
	PerformActionOnFolders(watchFolderVerbose, folderTree)
	commands := command.LoadCommands(ConfigurationData(), getPackageFolder())

	done := make(chan bool)
	go processEvent(commands)
	<-done

	err := watcher.Close()
	if err != nil {
		log.Fatal("Error watcher not properly closed:", err)
	}
}

func watchFolders(folders []string) {
	PerformActionOnFolders(watchFolder, folders)
}

func PerformActionOnFolders(action FolderAction, folders []string) {
	for _, folder := range folders {
		action(folder)
	}
}

func watchFolder(folder string) {
	error := watcher.Watch(folder)
	if error != nil {
		log.Fatal("Can't watch folder `", folder, "`. Error: ", error)
	}
}

func watchFolderVerbose(folder string) {
	log.Println("Watching '" + folder + "'.")
	watchFolder(folder)
}

func stopWatchFolders(folders []string) {
	PerformActionOnFolders(stopWatchFolder, folders)
}

func stopWatchFolder(folder string) {
	log.Println("NOT watching '" + folder + "'.")

	error := watcher.RemoveWatch(folder)
	if error != nil {
		log.Fatal("Can't stop watching folder `", folder, "`. Error: ", error)
	}
}

// Reads the event channel and takes the appropirate decision
func processEvent(commands []command.Command) {
	for {
		select {
		case ev := <-watcher.Event:
			if !ev.IsAttrib() {
				checkProjectHealth(commands)
			}
		case err := <-watcher.Error:
			log.Println("error:", err)
		}
	}
}

// Stops watching the files in the tree directory in case of
// any file modification triggered by the commands (infinite loop).
// Run all commands to check the project health
// Starts watching the files again
func checkProjectHealth(commands []command.Command) {
	stopWatchFolders(folderTree)
	command.RunCommands(commands)
	watchFolders(folderTree)
}
