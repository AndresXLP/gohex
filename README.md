# Gohex
Gohex is a CLI application that helps you create the initial files and folders to implement the hexagonal architecture in a microservice with Go.

---
# Intall
It can be installed by running:

`$ go install github.com/andresxlp/gohex@latest`

---
# Usage
1. Create a folder where you are going to work and enter it:
   - `$ mkdir myProyect && cd $_`
2. Execute gohex command:
   - `$ gohex new myProyect`
3. Set necessary env's to run  `SERVER_HOST` `SERVER_PORT`
   - `If you develop in IntelliJ or Goland IDE set env in your build configuration`
   - `If you develop in VSCode create an .env file in your root folder and check the config.go file`
    
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
