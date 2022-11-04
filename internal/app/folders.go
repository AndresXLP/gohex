package app

import (
	"os"

	"github.com/andresxlp/gohex/internal/enums"
)

func (fl *Service) CreateAllFolders() {
	createFolder(enums.CmdFolder)
	createFolder(enums.ProvidersFolder)
	createFolder(enums.ConfigFolder)
	createFolder(enums.InternalFolder)
	createFolder(enums.InfraFolder)
	createFolder(enums.ApiFolder)
	createFolder(enums.HandlerFolder)
	createFolder(enums.RouterFolder)
}

func createFolder(folder string) {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err = os.Mkdir(folder, 0755)
		if err != nil {
			panic(err)
		}
	}
}
