// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Aws
{
    /// <summary>
    /// Logs into a Vault server using an AWS auth backend. Login can be
    /// accomplished using a signed identity request from IAM or using ec2
    /// instance metadata. For more information, see the [Vault
    /// documentation](https://www.vaultproject.io/docs/auth/aws.html).
    /// </summary>
    [VaultResourceType("vault:aws/authBackendLogin:AuthBackendLogin")]
    public partial class AuthBackendLogin : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The token's accessor.
        /// </summary>
        [Output("accessor")]
        public Output<string> Accessor { get; private set; } = null!;

        /// <summary>
        /// The authentication type used to generate this token.
        /// </summary>
        [Output("authType")]
        public Output<string> AuthType { get; private set; } = null!;

        /// <summary>
        /// The unique name of the AWS auth backend. Defaults to
        /// 'aws'.
        /// </summary>
        [Output("backend")]
        public Output<string?> Backend { get; private set; } = null!;

        /// <summary>
        /// The token returned by Vault.
        /// </summary>
        [Output("clientToken")]
        public Output<string> ClientToken { get; private set; } = null!;

        /// <summary>
        /// The HTTP method used in the signed IAM
        /// request.
        /// </summary>
        [Output("iamHttpRequestMethod")]
        public Output<string?> IamHttpRequestMethod { get; private set; } = null!;

        /// <summary>
        /// The base64-encoded body of the signed
        /// request.
        /// </summary>
        [Output("iamRequestBody")]
        public Output<string?> IamRequestBody { get; private set; } = null!;

        /// <summary>
        /// The base64-encoded, JSON serialized
        /// representation of the GetCallerIdentity HTTP request headers.
        /// </summary>
        [Output("iamRequestHeaders")]
        public Output<string?> IamRequestHeaders { get; private set; } = null!;

        /// <summary>
        /// The base64-encoded HTTP URL used in the signed
        /// request.
        /// </summary>
        [Output("iamRequestUrl")]
        public Output<string?> IamRequestUrl { get; private set; } = null!;

        /// <summary>
        /// The base64-encoded EC2 instance identity document to
        /// authenticate with. Can be retrieved from the EC2 metadata server.
        /// </summary>
        [Output("identity")]
        public Output<string?> Identity { get; private set; } = null!;

        /// <summary>
        /// The duration in seconds the token will be valid, relative
        /// to the time in `lease_start_time`.
        /// </summary>
        [Output("leaseDuration")]
        public Output<int> LeaseDuration { get; private set; } = null!;

        [Output("leaseStartTime")]
        public Output<string> LeaseStartTime { get; private set; } = null!;

        /// <summary>
        /// A map of information returned by the Vault server about the
        /// authentication used to generate this token.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>> Metadata { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// The unique nonce to be used for login requests. Can be
        /// set to a user-specified value, or will contain the server-generated value
        /// once a token is issued. EC2 instances can only acquire a single token until
        /// the whitelist is tidied again unless they keep track of this nonce.
        /// </summary>
        [Output("nonce")]
        public Output<string> Nonce { get; private set; } = null!;

        /// <summary>
        /// The PKCS#7 signature of the identity document to
        /// authenticate with, with all newline characters removed. Can be retrieved from
        /// the EC2 metadata server.
        /// </summary>
        [Output("pkcs7")]
        public Output<string?> Pkcs7 { get; private set; } = null!;

        /// <summary>
        /// The Vault policies assigned to this token.
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<string>> Policies { get; private set; } = null!;

        /// <summary>
        /// Set to true if the token can be extended through renewal.
        /// </summary>
        [Output("renewable")]
        public Output<bool> Renewable { get; private set; } = null!;

        /// <summary>
        /// The name of the AWS auth backend role to create tokens
        /// against.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// The base64-encoded SHA256 RSA signature of the
        /// instance identity document to authenticate with, with all newline characters
        /// removed. Can be retrieved from the EC2 metadata server.
        /// </summary>
        [Output("signature")]
        public Output<string?> Signature { get; private set; } = null!;


        /// <summary>
        /// Create a AuthBackendLogin resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackendLogin(string name, AuthBackendLoginArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:aws/authBackendLogin:AuthBackendLogin", name, args ?? new AuthBackendLoginArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthBackendLogin(string name, Input<string> id, AuthBackendLoginState? state = null, CustomResourceOptions? options = null)
            : base("vault:aws/authBackendLogin:AuthBackendLogin", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "clientToken",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AuthBackendLogin resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthBackendLogin Get(string name, Input<string> id, AuthBackendLoginState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthBackendLogin(name, id, state, options);
        }
    }

    public sealed class AuthBackendLoginArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique name of the AWS auth backend. Defaults to
        /// 'aws'.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The HTTP method used in the signed IAM
        /// request.
        /// </summary>
        [Input("iamHttpRequestMethod")]
        public Input<string>? IamHttpRequestMethod { get; set; }

        /// <summary>
        /// The base64-encoded body of the signed
        /// request.
        /// </summary>
        [Input("iamRequestBody")]
        public Input<string>? IamRequestBody { get; set; }

        /// <summary>
        /// The base64-encoded, JSON serialized
        /// representation of the GetCallerIdentity HTTP request headers.
        /// </summary>
        [Input("iamRequestHeaders")]
        public Input<string>? IamRequestHeaders { get; set; }

        /// <summary>
        /// The base64-encoded HTTP URL used in the signed
        /// request.
        /// </summary>
        [Input("iamRequestUrl")]
        public Input<string>? IamRequestUrl { get; set; }

        /// <summary>
        /// The base64-encoded EC2 instance identity document to
        /// authenticate with. Can be retrieved from the EC2 metadata server.
        /// </summary>
        [Input("identity")]
        public Input<string>? Identity { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The unique nonce to be used for login requests. Can be
        /// set to a user-specified value, or will contain the server-generated value
        /// once a token is issued. EC2 instances can only acquire a single token until
        /// the whitelist is tidied again unless they keep track of this nonce.
        /// </summary>
        [Input("nonce")]
        public Input<string>? Nonce { get; set; }

        /// <summary>
        /// The PKCS#7 signature of the identity document to
        /// authenticate with, with all newline characters removed. Can be retrieved from
        /// the EC2 metadata server.
        /// </summary>
        [Input("pkcs7")]
        public Input<string>? Pkcs7 { get; set; }

        /// <summary>
        /// The name of the AWS auth backend role to create tokens
        /// against.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The base64-encoded SHA256 RSA signature of the
        /// instance identity document to authenticate with, with all newline characters
        /// removed. Can be retrieved from the EC2 metadata server.
        /// </summary>
        [Input("signature")]
        public Input<string>? Signature { get; set; }

        public AuthBackendLoginArgs()
        {
        }
        public static new AuthBackendLoginArgs Empty => new AuthBackendLoginArgs();
    }

    public sealed class AuthBackendLoginState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The token's accessor.
        /// </summary>
        [Input("accessor")]
        public Input<string>? Accessor { get; set; }

        /// <summary>
        /// The authentication type used to generate this token.
        /// </summary>
        [Input("authType")]
        public Input<string>? AuthType { get; set; }

        /// <summary>
        /// The unique name of the AWS auth backend. Defaults to
        /// 'aws'.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("clientToken")]
        private Input<string>? _clientToken;

        /// <summary>
        /// The token returned by Vault.
        /// </summary>
        public Input<string>? ClientToken
        {
            get => _clientToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The HTTP method used in the signed IAM
        /// request.
        /// </summary>
        [Input("iamHttpRequestMethod")]
        public Input<string>? IamHttpRequestMethod { get; set; }

        /// <summary>
        /// The base64-encoded body of the signed
        /// request.
        /// </summary>
        [Input("iamRequestBody")]
        public Input<string>? IamRequestBody { get; set; }

        /// <summary>
        /// The base64-encoded, JSON serialized
        /// representation of the GetCallerIdentity HTTP request headers.
        /// </summary>
        [Input("iamRequestHeaders")]
        public Input<string>? IamRequestHeaders { get; set; }

        /// <summary>
        /// The base64-encoded HTTP URL used in the signed
        /// request.
        /// </summary>
        [Input("iamRequestUrl")]
        public Input<string>? IamRequestUrl { get; set; }

        /// <summary>
        /// The base64-encoded EC2 instance identity document to
        /// authenticate with. Can be retrieved from the EC2 metadata server.
        /// </summary>
        [Input("identity")]
        public Input<string>? Identity { get; set; }

        /// <summary>
        /// The duration in seconds the token will be valid, relative
        /// to the time in `lease_start_time`.
        /// </summary>
        [Input("leaseDuration")]
        public Input<int>? LeaseDuration { get; set; }

        [Input("leaseStartTime")]
        public Input<string>? LeaseStartTime { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// A map of information returned by the Vault server about the
        /// authentication used to generate this token.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
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
        /// The unique nonce to be used for login requests. Can be
        /// set to a user-specified value, or will contain the server-generated value
        /// once a token is issued. EC2 instances can only acquire a single token until
        /// the whitelist is tidied again unless they keep track of this nonce.
        /// </summary>
        [Input("nonce")]
        public Input<string>? Nonce { get; set; }

        /// <summary>
        /// The PKCS#7 signature of the identity document to
        /// authenticate with, with all newline characters removed. Can be retrieved from
        /// the EC2 metadata server.
        /// </summary>
        [Input("pkcs7")]
        public Input<string>? Pkcs7 { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// The Vault policies assigned to this token.
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// Set to true if the token can be extended through renewal.
        /// </summary>
        [Input("renewable")]
        public Input<bool>? Renewable { get; set; }

        /// <summary>
        /// The name of the AWS auth backend role to create tokens
        /// against.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The base64-encoded SHA256 RSA signature of the
        /// instance identity document to authenticate with, with all newline characters
        /// removed. Can be retrieved from the EC2 metadata server.
        /// </summary>
        [Input("signature")]
        public Input<string>? Signature { get; set; }

        public AuthBackendLoginState()
        {
        }
        public static new AuthBackendLoginState Empty => new AuthBackendLoginState();
    }
}
