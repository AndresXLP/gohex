package enums

const (
	CmdFolder       = "cmd"
	ProvidersFolder = CmdFolder + "/providers"
	ConfigFolder    = "config"
	InternalFolder  = "internal"
	InfraFolder     = InternalFolder + "/infra"
	ApiFolder       = InfraFolder + "/api"
	HandlerFolder   = ApiFolder + "/handler"
	RouterFolder    = ApiFolder + "/router"
)
