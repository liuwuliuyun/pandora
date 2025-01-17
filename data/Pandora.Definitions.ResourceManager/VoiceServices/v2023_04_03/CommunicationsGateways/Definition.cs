using System.Collections.Generic;
using Pandora.Definitions.Interfaces;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.VoiceServices.v2023_04_03.CommunicationsGateways;

internal class Definition : ResourceDefinition
{
    public string Name => "CommunicationsGateways";
    public IEnumerable<Interfaces.ApiOperation> Operations => new List<Interfaces.ApiOperation>
    {
        new CreateOrUpdateOperation(),
        new DeleteOperation(),
        new GetOperation(),
        new ListByResourceGroupOperation(),
        new ListBySubscriptionOperation(),
        new UpdateOperation(),
    };
    public IEnumerable<System.Type> Constants => new List<System.Type>
    {
        typeof(AutoGeneratedDomainNameLabelScopeConstant),
        typeof(CommunicationsPlatformConstant),
        typeof(ConnectivityConstant),
        typeof(E911TypeConstant),
        typeof(ProvisioningStateConstant),
        typeof(StatusConstant),
        typeof(TeamsCodecsConstant),
    };
    public IEnumerable<System.Type> Models => new List<System.Type>
    {
        typeof(CommunicationsGatewayModel),
        typeof(CommunicationsGatewayPropertiesModel),
        typeof(CommunicationsGatewayUpdateModel),
        typeof(PrimaryRegionPropertiesModel),
        typeof(ServiceRegionPropertiesModel),
    };
}
