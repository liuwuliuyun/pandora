using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.DataBoxEdge.v2023_07_01.Users;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum UserTypeConstant
{
    [Description("ARM")]
    ARM,

    [Description("LocalManagement")]
    LocalManagement,

    [Description("Share")]
    Share,
}
