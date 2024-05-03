// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.TerraformCloud
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test = new Vault.TerraformCloud.SecretBackend("test", new()
    ///     {
    ///         Backend = "terraform",
    ///         Description = "Manages the Terraform Cloud backend",
    ///         Token = "V0idfhi2iksSDU234ucdbi2nidsi...",
    ///     });
    /// 
    ///     var example = new Vault.TerraformCloud.SecretRole("example", new()
    ///     {
    ///         Backend = test.Backend,
    ///         Name = "test-role",
    ///         Organization = "example-organization-name",
    ///         TeamId = "team-ieF4isC...",
    ///     });
    /// 
    ///     var token = new Vault.TerraformCloud.SecretCreds("token", new()
    ///     {
    ///         Backend = test.Backend,
    ///         Role = example.Name,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [VaultResourceType("vault:terraformcloud/secretCreds:SecretCreds")]
    public partial class SecretCreds : global::Pulumi.CustomResource
    {
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// The lease associated with the token. Only user tokens will have a 
        /// Vault lease associated with them.
        /// </summary>
        [Output("leaseId")]
        public Output<string> LeaseId { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// The organization associated with the token provided.
        /// </summary>
        [Output("organization")]
        public Output<string> Organization { get; private set; } = null!;

        /// <summary>
        /// Name of the role.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// The team id associated with the token provided.
        /// </summary>
        [Output("teamId")]
        public Output<string> TeamId { get; private set; } = null!;

        /// <summary>
        /// The actual token that was generated and can be used with API calls
        /// to identify the user of the call.
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        /// <summary>
        /// The public identifier for a specific token. It can be used 
        /// to look up information about a token or to revoke a token.
        /// </summary>
        [Output("tokenId")]
        public Output<string> TokenId { get; private set; } = null!;


        /// <summary>
        /// Create a SecretCreds resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretCreds(string name, SecretCredsArgs args, CustomResourceOptions? options = null)
            : base("vault:terraformcloud/secretCreds:SecretCreds", name, args ?? new SecretCredsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretCreds(string name, Input<string> id, SecretCredsState? state = null, CustomResourceOptions? options = null)
            : base("vault:terraformcloud/secretCreds:SecretCreds", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "leaseId",
                    "token",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecretCreds resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretCreds Get(string name, Input<string> id, SecretCredsState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretCreds(name, id, state, options);
        }
    }

    public sealed class SecretCredsArgs : global::Pulumi.ResourceArgs
    {
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Name of the role.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public SecretCredsArgs()
        {
        }
        public static new SecretCredsArgs Empty => new SecretCredsArgs();
    }

    public sealed class SecretCredsState : global::Pulumi.ResourceArgs
    {
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("leaseId")]
        private Input<string>? _leaseId;

        /// <summary>
        /// The lease associated with the token. Only user tokens will have a 
        /// Vault lease associated with them.
        /// </summary>
        public Input<string>? LeaseId
        {
            get => _leaseId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _leaseId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The organization associated with the token provided.
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        /// <summary>
        /// Name of the role.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The team id associated with the token provided.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        [Input("token")]
        private Input<string>? _token;

        /// <summary>
        /// The actual token that was generated and can be used with API calls
        /// to identify the user of the call.
        /// </summary>
        public Input<string>? Token
        {
            get => _token;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _token = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The public identifier for a specific token. It can be used 
        /// to look up information about a token or to revoke a token.
        /// </summary>
        [Input("tokenId")]
        public Input<string>? TokenId { get; set; }

        public SecretCredsState()
        {
        }
        public static new SecretCredsState Empty => new SecretCredsState();
    }
}
