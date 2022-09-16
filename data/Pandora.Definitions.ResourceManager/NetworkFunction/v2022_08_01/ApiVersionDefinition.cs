using System.Collections.Generic;
using Pandora.Definitions.Interfaces;

namespace Pandora.Definitions.ResourceManager.NetworkFunction.v2022_08_01;

public partial class Definition : ApiVersionDefinition
{
    public string ApiVersion => "2022-08-01";
    public bool Preview => false;
    public TransportLayer TransportLayer => TransportLayer.Autorest;
    public Source Source => Source.ResourceManagerRestApiSpecs;

    public IEnumerable<ResourceDefinition> Resources => new List<ResourceDefinition>
    {
        new AzureTrafficCollectors.Definition(),
        new CollectorPolicies.Definition(),
    };
}
