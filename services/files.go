package services

import (
	"fmt"
	"io/ioutil"

	"github.com/andresxlp/gohex/enums"
	"github.com/andresxlp/gohex/utils/templatesExec"
)

func CreateAllFiles(path, module string) {
	createMainFile(module)
	createDiFile(module)
	createConfigFile(module)
	createHealthFile()
	createRouterFile(module, path)
}

func createMainFile(module string) {
	main := fmt.Sprintf(enums.Separator, enums.CmdFolder, enums.MainFile)

	createFile(main, templatesExec.GetTemplateWhitValues(enums.GetMainFile, dataTemplate(module)))
}

func createDiFile(module string) {
	di := fmt.Sprintf(enums.Separator, enums.ProvidersFolder, enums.DiFile)

	createFile(di, templatesExec.GetTemplateWhitValues(enums.GetDiFile, dataTemplate(module)))
}

func createConfigFile(module string) {
	config := fmt.Sprintf(enums.Separator, enums.ConfigFolder, enums.ConfigFile)

	createFile(config, templatesExec.GetTemplateWhitValues(enums.GetConfigFile, dataTemplate(module)))
}

func createHealthFile() {
	health := fmt.Sprintf(enums.Separator, enums.HandlerFolder, enums.HealthFile)
	createFile(health, templatesExec.GetTemplateWhitValues(enums.GetHealthFile, map[string]interface{}{}))
}

func createRouterFile(module, path string) {
	router := fmt.Sprintf(enums.Separator, enums.RouterFolder, enums.RouterFile)

	createFile(router, templatesExec.GetTemplateWhitValues(enums.GetRouterFile, dataTemplate(module, path)))
}

func dataTemplate(data ...string) map[string]interface{} {
	var dataMap = map[string]interface{}{}

	switch len(data) {
	case 1:
		dataMap[enums.Module] = data[0]
		break
	case 2:
		dataMap[enums.Module] = data[0]
		dataMap[enums.BasePath] = data[1]
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
