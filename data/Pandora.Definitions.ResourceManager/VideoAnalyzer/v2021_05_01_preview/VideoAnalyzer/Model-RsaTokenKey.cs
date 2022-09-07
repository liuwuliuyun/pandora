using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.VideoAnalyzer.v2021_05_01_preview.VideoAnalyzer;

[ValueForType("#Microsoft.VideoAnalyzer.RsaTokenKey")]
internal class RsaTokenKeyModel : TokenKeyModel
{
    [JsonPropertyName("alg")]
    [Required]
    public AccessPolicyRsaAlgoConstant Alg { get; set; }

    [JsonPropertyName("e")]
    [Required]
    public string E { get; set; }

    [JsonPropertyName("n")]
    [Required]
    public string N { get; set; }
}