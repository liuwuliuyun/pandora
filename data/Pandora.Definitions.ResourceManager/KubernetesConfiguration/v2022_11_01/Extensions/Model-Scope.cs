using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.KubernetesConfiguration.v2022_11_01.Extensions;


internal class ScopeModel
{
    [JsonPropertyName("cluster")]
    public ScopeClusterModel? Cluster { get; set; }

    [JsonPropertyName("namespace")]
    public ScopeNamespaceModel? Namespace { get; set; }
}