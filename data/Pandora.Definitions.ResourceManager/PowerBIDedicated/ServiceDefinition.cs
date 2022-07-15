using Pandora.Definitions.Interfaces;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.PowerBIDedicated;

public partial class Service : ServiceDefinition
{
    public string Name => "PowerBIDedicated";
    public string? ResourceProvider => "Microsoft.PowerBIdedicated";
    public string? TerraformPackageName => "powerbidedicated";
}
