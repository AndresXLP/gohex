package templatesExec

import (
	"strings"
	"sync"
	"text/template"

	"github.com/andresxlp/gohex/enums"
	templ "github.com/andresxlp/gohex/utils/templates"
)

var (
	once      sync.Once
	templates map[enums.TemplateLabel]string
)

func TemplateMust(value enums.TemplateLabel, sql string) *template.Template {
	return template.Must(template.New(value.String()).Parse(sql))
}

func ExecuteTemplate(templ *template.Template, data map[string]interface{}) string {
	builder := &strings.Builder{}
	if err := templ.Execute(builder, data); err != nil {
		panic(err)
	}

	return builder.String()
}

func GetTemplateWhitValues(value enums.TemplateLabel, data map[string]interface{}) string {
	tmp := TemplateMust(value, getTemplate(value))
	return ExecuteTemplate(tmp, data)
}

func getTemplate(value enums.TemplateLabel) string {
	once.Do(func() {
		templates = map[enums.TemplateLabel]string{
			enums.GetMainFile:   templ.MainFile,
			enums.GetDiFile:     templ.DiFile,
			enums.GetConfigFile: templ.ConfigFile,
			enums.GetHealthFile: templ.HealthFile,
			enums.GetRouterFile: templ.RouterFile,
		}
	})

	return templates[value]
}
