using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.LoadTestService.v2022_12_01.Quotas;


internal class QuotaResourcePropertiesModel
{
    [JsonPropertyName("limit")]
    public int? Limit { get; set; }

    [JsonPropertyName("provisioningState")]
    public ResourceStateConstant? ProvisioningState { get; set; }

    [JsonPropertyName("usage")]
    public int? Usage { get; set; }
}
