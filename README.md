# toggl CLI client

[toggl](https://toggl.com/) CLI Client, written in Golang.

## Description

[toggl](https://toggl.com/) is a time tracking web application.
This program will let you use the toggl in CLI.

## Usage

![demo image](https://cloud.githubusercontent.com/assets/6121271/21588531/0108bd18-d12b-11e6-9fdc-e65aa1f15768.gif)

```
$ toggl --help
NAME:
   toggl - Toggl API CLI Client

USAGE:
   toggl [global options] command [command options] [arguments...]

VERSION:
   0.4.0

AUTHOR(S):
   sachaos <sakataku7@gmail.com>

COMMANDS:
     start       Start time entry
     stop        End time entry
     current     Show current time entry
     workspaces  Show workspaces
     projects    Show projects on current workspaces
     local       Set current dir workspace
     global      Set global workspace
     help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --cache
   --help, -h     show help
   --version, -v  print the version
```

## Install

To install, use `go get`:

```bash
$ go get github.com/sachaos/toggl
```

### Homebrew

```shell
$ brew install sachaos/tap/toggl
```

### Register API token

When you run `toggl` first time, you will be asked your toggl API token.
Please input toggl API token and register it.

## Author

[sachaos](https://github.com/sachaos)
