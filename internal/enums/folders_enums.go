package enums

const (
	CmdFolder        = "cmd"
	ProvidersFolder  = CmdFolder + "/providers"
	ConfigFolder     = "config"
	InternalFolder   = "internal"
	AppFolder        = InternalFolder + "/app"
	ConstantsFolder  = InternalFolder + "/constants"
	DomainFolder     = InternalFolder + "/domain"
	DTOFolder        = DomainFolder + "/dto"
	EntityFolder     = DomainFolder + "/entity"
	PortsFolder      = DomainFolder + "/ports"
	InfraFolder      = InternalFolder + "/infra"
	AdaptersFolder   = InfraFolder + "/adapters"
	ApiFolder        = InfraFolder + "/api"
	ResourcesFolder  = InfraFolder + "/resource"
	HandlerFolder    = ApiFolder + "/handler"
	MiddlewareFolder = ApiFolder + "/middleware"
	RouterFolder     = ApiFolder + "/router"
	GroupsFolder     = RouterFolder + "/groups"
)
