using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.ManagementGroups.v2023_04_01.ManagementGroups;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum ManagementGroupExpandTypeConstant
{
    [Description("ancestors")]
    Ancestors,

    [Description("children")]
    Children,

    [Description("path")]
    Path,
}
