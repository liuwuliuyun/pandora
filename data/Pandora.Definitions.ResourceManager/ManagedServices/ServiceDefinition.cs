using Pandora.Definitions.Interfaces;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.ManagedServices;

public partial class Service : ServiceDefinition
{
    public string Name => "ManagedServices";
    public string? ResourceProvider => "Microsoft.ManagedServices";
    public string? TerraformPackageName => "managedservices";
}
