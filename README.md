# Gohex
Gohex is a CLI application that helps you create the initial files and folders to implement the hexagonal architecture in a microservice with Go.

---
# Intall
It can be installed by running:

`$ go install github.com/andresxlp/gohex@latest`

---
# Usage
Create the folder where you are going to host the repository,
we recommend that you use your $GOPATH

![image](https://github.com/AndresXLP/gohex/blob/main/gohex-usage.png)

---

# Structure
This generator has basic structure like this
```
Folder Struct:
  MS FOLDER
  ├─ cmd
  │  └─ providers
  ├─ config
  └─ internal
     └─ infra
        └─ api
           ├─ handler
           └─ router
```

---
# Dependency
As main dependencies for the creation of the initial structure we use:
- [echo](https://echo.labstack.com/) - `High performance, extensible, minimalist Go web framework.`
- [dig](https://github.com/uber-go/dig) - `A reflection based dependency injection toolkit for Go.`
- [envconfig](http://github.com/kelseyhightower/envconfig) - `Decoding of environment variables.`

---
# Environments
When Gohex successfully creates the initial structure, don't forget set the environment variables to run your Microservice locally

* `SERVER_HOST`: host for the server
* `SERVER_PORT:` port for the server

---
# Creator
- [Andres Puello](https://github.com/AndresXLP/gohex)

---
# Contributors