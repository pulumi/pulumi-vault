// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Database.Outputs
{

    [OutputType]
    public sealed class SecretsMountMongodbatla
    {
        /// <summary>
        /// A list of roles that are allowed to use this
        /// connection.
        /// </summary>
        public readonly ImmutableArray<string> AllowedRoles;
        /// <summary>
        /// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
        /// 
        /// Supported list of database secrets engines that can be configured:
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Data;
        public readonly string Name;
        /// <summary>
        /// Specifies the name of the plugin to use.
        /// </summary>
        public readonly string? PluginName;
        /// <summary>
        /// The Private Programmatic API Key used to connect with MongoDB Atlas API.
        /// </summary>
        public readonly string PrivateKey;
        /// <summary>
        /// The Project ID the Database User should be created within.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// The Public Programmatic API Key used to authenticate with the MongoDB Atlas API.
        /// </summary>
        public readonly string PublicKey;
        /// <summary>
        /// A list of database statements to be executed to rotate the root user's credentials.
        /// </summary>
        public readonly ImmutableArray<string> RootRotationStatements;
        /// <summary>
        /// Whether the connection should be verified on
        /// initial configuration or not.
        /// </summary>
        public readonly bool? VerifyConnection;

        [OutputConstructor]
        private SecretsMountMongodbatla(
            ImmutableArray<string> allowedRoles,

            ImmutableDictionary<string, object>? data,

            string name,

            string? pluginName,

            string privateKey,

            string projectId,

            string publicKey,

            ImmutableArray<string> rootRotationStatements,

            bool? verifyConnection)
        {
            AllowedRoles = allowedRoles;
            Data = data;
            Name = name;
            PluginName = pluginName;
            PrivateKey = privateKey;
            ProjectId = projectId;
            PublicKey = publicKey;
            RootRotationStatements = rootRotationStatements;
            VerifyConnection = verifyConnection;
        }
    }
}
