using System.Collections.Generic;
using Pandora.Definitions.Interfaces;

namespace Pandora.Definitions.TestData.Pandamonium.v2020_01_01;

public class Definition : ApiVersionDefinition
{
    public string ApiVersion => "2020-01-01";
    public bool Generate => true;
    public bool Preview => false;
	public string TransportLayer => "pandora";

    public IEnumerable<ResourceDefinition> Resources => new List<ResourceDefinition>
    {
        new Grouping.Definition()
    };

    public Source Source => Source.HandWritten;
}