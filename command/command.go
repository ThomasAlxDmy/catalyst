package command

import (
  "fmt"
  "log"
  "strings"
  "os/exec"

  "gopkg.in/yaml.v1"
)

var (
  packagePath = ""
  )

// Pretty structure to extract the command from the configuration file
type CommandsConfig struct{
  Commands []Command
}

// A command is a simple structure that contains a name and a list of arguments
type Command struct{
  Name string
  Arguments []string
}

// Runs a command with it arguments
// Prints the results only if there's an output or if there's an error
func (command *Command) Run() error {
  cmd := exec.Command(command.Name, command.Arguments...)

  output, err := cmd.CombinedOutput()
  commandOptionsStr := strings.Join(command.Arguments, " ")
  result := "Running command `" + command.Name + " " + fmt.Sprintf("%s", commandOptionsStr) + "`"
  result += ". Output: \n" + fmt.Sprintf("%s", output)

  if err != nil {
    log.Println(result, "error:", err)
  } else if len(output) > 0 {
    log.Println(result)
  }

  return err
}

// Replaces each argument with the value `$1` with the packagePath if packagePath has been given
func (command *Command) sanitizeArguments(){
  if packagePath != "" {
    for i, argument := range command.Arguments{
      if argument == "$1" {
        command.Arguments[i] = packagePath
      }
    }
  }
}

func RunCommands(commands []Command){
  for _, command := range commands {
    err := command.Run()
    if err != nil {
      break
    }
  }
}

// Loads command from the yaml file and sanitized them
func LoadCommands(data []byte, packageFolder string) []Command{
  packagePath = packageFolder

  commands := extractCommandFromConfig(data)
  for _, command := range commands {
    command.sanitizeArguments()
  }

  return commands
}

// Loads command from the yaml
func extractCommandFromConfig(data []byte) []Command{
  commandsConfig := CommandsConfig{}
  err := yaml.Unmarshal(data, &commandsConfig)
  if err != nil {
    log.Fatalf("Error unmarshaling yaml config: %v", err)
  }

  return commandsConfig.Commands
}
