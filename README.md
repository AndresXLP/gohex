# gohex
Gohex is a CLI application that helps you create the initial files and folders to implement the hexagonal architecture in a microservice with Go.

# Usage
It can be installed by running:

`go install github.com/andresxlp/gohexa@latest`

```
Usage:
  gohex [command]

Examples:
  MS FOLDER
  ├─ cmd
  │  └─ providers
  ├─ config
  └─ internal
     └─ infra
        └─ api
           ├─ handler
           └─ router

Available Commands:
  init        Initializes the creation of the files and folders.

Flags:
  -h, --help   help for gohex

Use "gohex [command] --help" for more information about a command.
```
