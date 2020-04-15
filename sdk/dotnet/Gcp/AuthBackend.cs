// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Gcp
{
    /// <summary>
    /// Provides a resource to configure the [GCP auth backend within Vault](https://www.vaultproject.io/docs/auth/gcp.html).
    /// </summary>
    public partial class AuthBackend : Pulumi.CustomResource
    {
        /// <summary>
        /// The clients email associated with the credentials
        /// </summary>
        [Output("clientEmail")]
        public Output<string> ClientEmail { get; private set; } = null!;

        /// <summary>
        /// The Client ID of the credentials
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
        /// </summary>
        [Output("credentials")]
        public Output<string?> Credentials { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        /// <summary>
        /// The ID of the private key from the credentials
        /// </summary>
        [Output("privateKeyId")]
        public Output<string> PrivateKeyId { get; private set; } = null!;

        /// <summary>
        /// The GCP Project ID
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;


        /// <summary>
        /// Create a AuthBackend resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackend(string name, AuthBackendArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:gcp/authBackend:AuthBackend", name, args ?? new AuthBackendArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthBackend(string name, Input<string> id, AuthBackendState? state = null, CustomResourceOptions? options = null)
            : base("vault:gcp/authBackend:AuthBackend", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AuthBackend resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthBackend Get(string name, Input<string> id, AuthBackendState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthBackend(name, id, state, options);
        }
    }

    public sealed class AuthBackendArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The clients email associated with the credentials
        /// </summary>
        [Input("clientEmail")]
        public Input<string>? ClientEmail { get; set; }

        /// <summary>
        /// The Client ID of the credentials
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
        /// </summary>
        [Input("credentials")]
        public Input<string>? Credentials { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The ID of the private key from the credentials
        /// </summary>
        [Input("privateKeyId")]
        public Input<string>? PrivateKeyId { get; set; }

        /// <summary>
        /// The GCP Project ID
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public AuthBackendArgs()
        {
        }
    }

    public sealed class AuthBackendState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The clients email associated with the credentials
        /// </summary>
        [Input("clientEmail")]
        public Input<string>? ClientEmail { get; set; }

        /// <summary>
        /// The Client ID of the credentials
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
        /// </summary>
        [Input("credentials")]
        public Input<string>? Credentials { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The ID of the private key from the credentials
        /// </summary>
        [Input("privateKeyId")]
        public Input<string>? PrivateKeyId { get; set; }

        /// <summary>
        /// The GCP Project ID
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public AuthBackendState()
        {
        }
    }
}
