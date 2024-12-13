// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Azure.Inputs
{

    public sealed class BackendRoleAzureRoleArgs : global::Pulumi.ResourceArgs
    {
        [Input("roleId")]
        public Input<string>? RoleId { get; set; }

        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        public BackendRoleAzureRoleArgs()
        {
        }
        public static new BackendRoleAzureRoleArgs Empty => new BackendRoleAzureRoleArgs();
    }
}