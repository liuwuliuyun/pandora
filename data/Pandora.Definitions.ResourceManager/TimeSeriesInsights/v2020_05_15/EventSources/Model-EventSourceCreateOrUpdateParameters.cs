using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.TimeSeriesInsights.v2020_05_15.EventSources;


internal abstract class EventSourceCreateOrUpdateParametersModel
{
    [JsonPropertyName("kind")]
    [ProvidesTypeHint]
    [Required]
    public EventSourceKindConstant Kind { get; set; }

    [JsonPropertyName("localTimestamp")]
    public LocalTimestampModel? LocalTimestamp { get; set; }

    [JsonPropertyName("location")]
    [Required]
    public CustomTypes.Location Location { get; set; }

    [JsonPropertyName("tags")]
    public CustomTypes.Tags? Tags { get; set; }
}
