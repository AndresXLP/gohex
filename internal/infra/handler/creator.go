package handler

import (
	"log"
	"os/exec"
	"strings"

	apps "github.com/andresxlp/gohex/internal/app"
)

var (
	app *apps.Service
)

type AnyError struct {
	Err error
}

func CreateFolderAndFile(module string) {
	pathModule := getModule(module)

	go func() {
		app.CreateAllFolders()
		app.CreateAllFiles(pathModule)
	}()
	app.ProgressBarFilesAndFolders()
}

func ExecuteGoModule(module string) (r AnyError) {
	outputChannel := make(chan AnyError)
	go func() {
		if err := app.InitGoModule(module); err != nil {
			outputChannel <- AnyError{err}
		} else {
			outputChannel <- AnyError{nil}
		}
	}()
	app.ProgressBarGoModules()
	return <-outputChannel
}

func getModule(module string) string {
	var pathModule string
	if module != "" {
		return module
	}

	bytes, err := exec.Command("pwd").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	str := string(bytes)
	sliceString := strings.Split(str, "/")
	var pwd []string
	for i, s := range sliceString {
		if s == "src" {
			pwd = sliceString[i+1:]
			break
		}
	}
	pathModule = strings.Replace(strings.Join(pwd, "/"), "\n", "", 1)

	return pathModule
}
