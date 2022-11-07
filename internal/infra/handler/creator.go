package handler

import (
	apps "github.com/andresxlp/gohex/internal/app"
)

var (
	app      *apps.Service
	BasePath string
	Module   string
)

type AnyError struct {
	Err error
}

func VerifyFilesAndFolderExisting() (r AnyError) {
	outputChannel := make(chan AnyError)
	go func() {
		if err := app.FileExisting(); err != nil {
			outputChannel <- AnyError{err}
		} else {
			outputChannel <- AnyError{nil}
		}
	}()
	app.ProgressBarVerifyFilesAndFolder()
	return <-outputChannel
}

func CreateFolderAndFile() {
	go func() {
		app.CreateAllFolders()
		app.CreateAllFiles(Module, BasePath)
	}()
	app.ProgressBarFilesAndFolders()
}

func ExecuteGoModule() (r AnyError) {
	outputChannel := make(chan AnyError)
	go func() {
		if err := app.InitGoModule(); err != nil {
			outputChannel <- AnyError{err}
		} else {
			outputChannel <- AnyError{nil}
		}
	}()
	app.ProgressBarGoModules()
	return <-outputChannel
}
