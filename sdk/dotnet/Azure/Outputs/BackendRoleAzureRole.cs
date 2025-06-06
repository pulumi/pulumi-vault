// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Azure.Outputs
{

    [OutputType]
    public sealed class BackendRoleAzureRole
    {
        public readonly string? RoleId;
        public readonly string? RoleName;
        public readonly string Scope;

        [OutputConstructor]
        private BackendRoleAzureRole(
            string? roleId,

            string? roleName,

            string scope)
        {
            RoleId = roleId;
            RoleName = roleName;
            Scope = scope;
        }
    }
}
