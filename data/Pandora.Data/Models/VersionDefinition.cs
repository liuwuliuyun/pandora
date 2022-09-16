using System.Collections.Generic;

namespace Pandora.Data.Models;

public class VersionDefinition
{
    public string Version { get; set; }

    public bool Generate { get; set; }

    public bool Preview { get; set; }

    public ApiDefinitionsTransportLayer ApiDefinitionsTransportLayer { get; set; }

    public IEnumerable<ResourceDefinition> Resources { get; set; }

    public ApiDefinitionsSource Source { get; set; }
}