using System.Collections.Generic;
using Pandora.Definitions.Interfaces;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.PostgreSqlHSC.v2020_10_05_privatepreview.Configurations;

internal class Definition : ResourceDefinition
{
    public string Name => "Configurations";
    public IEnumerable<Interfaces.ApiOperation> Operations => new List<Interfaces.ApiOperation>
    {
        new GetOperation(),
        new ListByServerOperation(),
        new ListByServerGroupOperation(),
        new UpdateOperation(),
    };
    public IEnumerable<System.Type> Constants => new List<System.Type>
    {
        typeof(ConfigurationDataTypeConstant),
        typeof(ServerRoleConstant),
    };
    public IEnumerable<System.Type> Models => new List<System.Type>
    {
        typeof(ServerConfigurationModel),
        typeof(ServerConfigurationPropertiesModel),
        typeof(ServerGroupConfigurationModel),
        typeof(ServerGroupConfigurationPropertiesModel),
        typeof(ServerRoleGroupConfigurationModel),
    };
}
