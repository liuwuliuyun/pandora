using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.Sql.v2022_11_01_preview.ManagedDatabaseSensitivityLabels;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum RecommendedSensitivityLabelUpdateKindConstant
{
    [Description("disable")]
    Disable,

    [Description("enable")]
    Enable,
}
