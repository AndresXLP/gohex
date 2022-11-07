package app

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/andresxlp/gohex/internal/enums"
	"github.com/andresxlp/gohex/internal/utils/templatesExec"
)

type Service struct{}

var (
	files = []string{enums.MainFile, enums.DiFile, enums.ConfigFile, enums.HealthFile, enums.RouterFile, enums.GoModFile, enums.GoSumFile}
)

func (s *Service) CreateAllFiles(module, basePath string) {
	for _, file := range files[:5] {
		createFiles(file, module, basePath)
	}
}

func (s *Service) FileExisting() error {
	for _, file := range files {
		if verifyExistingFile(file) {
			return errors.New("The File " + file + " already exist.")
		}
	}
	return nil
}

func createFiles(fileName, module, basePath string) {
	queryTempl := map[string]enums.TemplateLabel{
		files[0]: enums.GetMainFile,
		files[1]: enums.GetDiFile,
		files[2]: enums.GetConfigFile,
		files[3]: enums.GetHealthFile,
		files[4]: enums.GetRouterFile,
	}

	createFile(fileName, templatesExec.GetTemplateWhitValues(queryTempl[fileName], dataTemplate(module, basePath)))
}

func dataTemplate(module, basePath string) map[string]interface{} {
	dataMap := map[string]interface{}{
		enums.Module:   module,
		enums.BasePath: basePath,
	}
	return dataMap
}

func createFile(fileName string, data string) {
	err := ioutil.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		panic(err)
	}
}

func verifyExistingFile(fileName string) bool {
	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		return true
	}
	return false
}
