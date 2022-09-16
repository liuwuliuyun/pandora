package discovery

import "github.com/hashicorp/pandora/tools/sdk/config/definitions"

type ServiceInput struct {
	RootNamespace              string
	ServiceName                string
	ApiVersion                 string
	OutputDirectory            string
	ResourceProvider           *string
	SwaggerDirectory           string
	SwaggerFiles               []string
	TerraformServiceDefinition *definitions.ServiceDefinition
	Transport                  Transport
}

type Transport = string

const (
	TransportAutorest Transport = "autorest"
	TransportPandora  Transport = "pandora"
)

type ResourceManagerServiceInput struct {
	ServiceName                string
	ApiVersion                 string
	OutputDirectory            string
	ResourceProvider           string
	SwaggerDirectory           string
	SwaggerFiles               []string
	TerraformServiceDefinition *definitions.ServiceDefinition
	Transport                  Transport
}

func (rmi ResourceManagerServiceInput) ToRunInput() ServiceInput {
	return ServiceInput{
		RootNamespace:              "Pandora.Definitions.ResourceManager",
		ServiceName:                rmi.ServiceName,
		ApiVersion:                 rmi.ApiVersion,
		ResourceProvider:           &rmi.ResourceProvider,
		OutputDirectory:            rmi.OutputDirectory,
		SwaggerDirectory:           rmi.SwaggerDirectory,
		SwaggerFiles:               rmi.SwaggerFiles,
		TerraformServiceDefinition: rmi.TerraformServiceDefinition,
		Transport:                  rmi.Transport,
	}
}
