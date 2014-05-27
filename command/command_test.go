package command

import (
  "os"
  "path"
  "reflect"
  "testing"
  )

func TestExtractCommandFromConfig(t *testing.T) {
  commands := extractCommandFromConfig(simpleYamlData)

  if !reflect.DeepEqual(commands, commandFromSimpleYaml) {
    t.Fatalf("Error unexpected command found. Got:", commands, "expected:", commandFromSimpleYaml)
  }
}

func TestLoadCommands(t *testing.T) {
  folder, err := os.Getwd()
  if err != nil {
    panic(err)
  }

  commands := LoadCommands(simpleYamlData, path.Dir(folder))

  if !reflect.DeepEqual(commands, commandFromSimpleYaml) {
    t.Fatalf("Error unexpected command found. Got:", commands, "expected:", commandFromSimpleYaml)
  }
}

func TestSanitizeArguments(t *testing.T) {
  TestSanitizeArgumentsSimpleYaml(t)
  TestSanitizeArgumentsComplexYaml(t)
}

func TestSanitizeArgumentsSimpleYaml(t *testing.T) {
  simpleCommands := extractCommandFromConfig(simpleYamlData)
  virginSimpleCommands := extractCommandFromConfig(simpleYamlData)

  for _, command := range extractCommandFromConfig(simpleYamlData) {
    command.sanitizeArguments()
  }

  if !reflect.DeepEqual(simpleCommands, virginSimpleCommands) {
    t.Fatalf("Error unexpected command found. Got:", simpleCommands, "expected:", virginSimpleCommands)
  }
}

func TestSanitizeArgumentsComplexYaml(t *testing.T){
  for _, command := range extractCommandFromConfig(simpleYamlData) {
    command.sanitizeArguments()

    for _, argument := range command.Arguments {
      if argument == "$1" {
        t.Fatalf("Arguments not sanatized properly:", argument)
      }
    }
  }
}

var simpleYamlData = []byte{ 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x3a, 0xa, 0x20, 0x20, 0x2d, 0x20,
                       0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x20, 0x63, 0x6c, 0x65, 0x61, 0x72, 0xa, 0x20, 0x20,
                       0x2d, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x20, 0x74, 0x6f, 0x75, 0x63, 0x68, 0xa,
                       0x20, 0x20, 0x20, 0x20, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3a,
                       0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x20, 0x74, 0x6d, 0x70,
                     }

var commandFromSimpleYaml = []Command{
                          Command{ Name:"clear", Arguments:[]string(nil)},
                          Command{ Name:"touch", Arguments:[]string{"tmp"}},
                        }

var complexYamlData = []byte{ 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x3a, 0xa, 0x20, 0x20, 0x2d, 0x20, 0x6e,
                        0x61, 0x6d, 0x65, 0x3a, 0x20, 0x63, 0x6c, 0x65, 0x61, 0x72, 0xa, 0x20, 0x20, 0x2d, 0x20,
                        0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x20, 0x67, 0x6f, 0xa, 0x20, 0x20, 0x20, 0x20, 0x61, 0x72,
                        0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
                        0x2d, 0x20, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20,
                        0x20, 0x2d, 0x20, 0x24, 0x31, 0xa, 0x20, 0x20, 0x2d, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3a,
                        0x20, 0x67, 0x6f, 0xa, 0x20, 0x20, 0x20, 0x20, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e,
                        0x74, 0x73, 0x3a, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x20, 0x66, 0x6d, 0x74,
                        0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x20, 0x24, 0x31, 0xa, 0x20, 0x20, 0x2d,
                        0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x20, 0x67, 0x6f, 0xa, 0x20, 0x20, 0x20, 0x20, 0x61,
                        0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20,
                        0x20, 0x2d, 0x20, 0x74, 0x65, 0x73, 0x74, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2d,
                        0x20, 0x2e, 0x2f, 0x2e, 0x2e, 0x2e,
                      }

var commandFromComplexYaml = []Command{
                                Command {Name:"clear", Arguments:[]string(nil)},
                                Command {Name:"go", Arguments:[]string{"install", "$1"}},
                                Command {Name:"go", Arguments:[]string{"fmt", "$1"}},
                                Command {Name:"go", Arguments:[]string{"test", "./..."}}}
