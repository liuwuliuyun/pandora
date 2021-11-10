using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.DataLakeStore.v2016_11_01.Accounts
{
    [ConstantType(ConstantTypeAttribute.ConstantType.String)]
    internal enum EncryptionStateConstant
    {
        [Description("Disabled")]
        Disabled,

        [Description("Enabled")]
        Enabled,
    }
}
