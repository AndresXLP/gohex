package enums

type TemplateLabel int

const (
	GetMainFile TemplateLabel = iota
	GetDiFile
	GetConfigFile
	GetHealthFile
	GetRouterFile
)

var templateLabel = [...]string{
	"GetMainFile",
	"GetDiFile",
	"GetConfigFile",
	"GetHealthFile",
	"GetRouterFile",
}

func (value TemplateLabel) String() string {
	return templateLabel[value]
}
