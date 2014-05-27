# Catalyst
========

Catalyst is a program created to considerably speed up your development environment.
Run it at the base of a go project and it will monitor all files present in it, and run commands each time a file was modifed.

## Install

  go get github.com/ThomasAlxDmy/Catalyst

## How it works?

Catalyst is a simple monitoring file system that runs given commands after each file modification (creation and deletion included).
At start up catalyst loads the commands from a yaml file into memory. By default, Catalyst uses `github.com/ThomasAlxDmy/Catalyst/catalyst_config` as configuration file, with the following commands:

  * clear
  * go install currentPackage
  * go test ./...

Because different projects require different needs, you can use a custom configuration. In order to do so, just create a new file called catalyst_config.yml at the root of your go project. Be carefull to respect the correct syntax and structure.

## Use

To use it, just go at the base folder of a go project and run `catalyst`. At first it will tell you files that are being monitored and will wait for you to perform an action.

