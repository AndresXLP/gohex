package app

import (
	"os"

	"github.com/andresxlp/gohex/internal/enums"
)

var folders = []string{
	enums.CmdFolder,
	enums.ProvidersFolder,
	enums.ConfigFolder,
	enums.InternalFolder,
	enums.InfraFolder,
	enums.ApiFolder,
	enums.HandlerFolder,
	enums.RouterFolder,
}

func (s *Service) CreateAllFolders() {
	for _, folder := range folders {
		createFolder(folder)
	}
}

func createFolder(folder string) {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err = os.Mkdir(folder, 0755)
		if err != nil {
			panic(err)
		}
	}
}
