package app

import (
	"fmt"
	"io/ioutil"

	enums2 "github.com/andresxlp/gohex/internal/enums"
	"github.com/andresxlp/gohex/internal/utils/templatesExec"
)

type Service struct{}

func (fl *Service) CreateAllFiles(module, basePath string) {
	createMainFile(module)
	createDiFile(module)
	createConfigFile(module)
	createHealthFile()
	createRouterFile(module, basePath)
}

func createMainFile(module string) {
	main := fmt.Sprintf(enums2.Separator, enums2.CmdFolder, enums2.MainFile)
	createFile(main, templatesExec.GetTemplateWhitValues(enums2.GetMainFile, dataTemplate(module)))
}

func createDiFile(module string) {
	di := fmt.Sprintf(enums2.Separator, enums2.ProvidersFolder, enums2.DiFile)

	createFile(di, templatesExec.GetTemplateWhitValues(enums2.GetDiFile, dataTemplate(module)))
}

func createConfigFile(module string) {
	config := fmt.Sprintf(enums2.Separator, enums2.ConfigFolder, enums2.ConfigFile)

	createFile(config, templatesExec.GetTemplateWhitValues(enums2.GetConfigFile, dataTemplate(module)))
}

func createHealthFile() {
	health := fmt.Sprintf(enums2.Separator, enums2.HandlerFolder, enums2.HealthFile)
	createFile(health, templatesExec.GetTemplateWhitValues(enums2.GetHealthFile, map[string]interface{}{}))
}

func createRouterFile(module, basePath string) {
	router := fmt.Sprintf(enums2.Separator, enums2.RouterFolder, enums2.RouterFile)

	createFile(router, templatesExec.GetTemplateWhitValues(enums2.GetRouterFile, dataTemplate(module, basePath)))
}

func dataTemplate(data ...string) map[string]interface{} {
	var dataMap = map[string]interface{}{}

	switch len(data) {
	case 1:
		dataMap[enums2.Module] = data[0]
		break
	case 2:
		dataMap[enums2.Module] = data[0]
		dataMap[enums2.BasePath] = data[1]
		break
	}

	return dataMap
}

func createFile(rutaDestino string, data string) {
	err := ioutil.WriteFile(rutaDestino, []byte(data), 0644)
	if err != nil {
		panic(err)
	}
}
