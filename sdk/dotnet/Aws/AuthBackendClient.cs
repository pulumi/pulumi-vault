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
    /// ## Example Usage
    /// 
    /// You can setup the AWS auth engine with Workload Identity Federation (WIF) for a secret-less configuration:
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Vault.AuthBackend("example", new()
    ///     {
    ///         Type = "aws",
    ///     });
    /// 
    ///     var exampleAuthBackendClient = new Vault.Aws.AuthBackendClient("example", new()
    ///     {
    ///         IdentityTokenAudience = "&lt;TOKEN_AUDIENCE&gt;",
    ///         IdentityTokenTtl = "&lt;TOKEN_TTL&gt;",
    ///         RoleArn = "&lt;AWS_ROLE_ARN&gt;",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Vault.AuthBackend("example", new()
    ///     {
    ///         Type = "aws",
    ///     });
    /// 
    ///     var exampleAuthBackendClient = new Vault.Aws.AuthBackendClient("example", new()
    ///     {
    ///         Backend = example.Path,
    ///         AccessKey = "INSERT_AWS_ACCESS_KEY",
    ///         SecretKey = "INSERT_AWS_SECRET_KEY",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// AWS auth backend clients can be imported using `auth/`, the `backend` path, and `/config/client` e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:aws/authBackendClient:AuthBackendClient example auth/aws/config/client
    /// ```
    /// </summary>
    [VaultResourceType("vault:aws/authBackendClient:AuthBackendClient")]
    public partial class AuthBackendClient : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The AWS access key that Vault should use for the
        /// auth backend. Mutually exclusive with `identity_token_audience`.
        /// </summary>
        [Output("accessKey")]
        public Output<string?> AccessKey { get; private set; } = null!;

        /// <summary>
        /// The path the AWS auth backend being configured was
        /// mounted at.  Defaults to `aws`.
        /// </summary>
        [Output("backend")]
        public Output<string?> Backend { get; private set; } = null!;

        /// <summary>
        /// Override the URL Vault uses when making EC2 API
        /// calls.
        /// </summary>
        [Output("ec2Endpoint")]
        public Output<string?> Ec2Endpoint { get; private set; } = null!;

        /// <summary>
        /// Override the URL Vault uses when making IAM API
        /// calls.
        /// </summary>
        [Output("iamEndpoint")]
        public Output<string?> IamEndpoint { get; private set; } = null!;

        /// <summary>
        /// The value to require in the
        /// `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
        /// that are used in the IAM auth method.
        /// </summary>
        [Output("iamServerIdHeaderValue")]
        public Output<string?> IamServerIdHeaderValue { get; private set; } = null!;

        /// <summary>
        /// The audience claim value. Mutually exclusive with `access_key`. 
        /// Requires Vault 1.17+. *Available only for Vault Enterprise*
        /// </summary>
        [Output("identityTokenAudience")]
        public Output<string?> IdentityTokenAudience { get; private set; } = null!;

        /// <summary>
        /// The TTL of generated identity tokens in seconds. Requires Vault 1.17+.
        /// *Available only for Vault Enterprise*
        /// </summary>
        [Output("identityTokenTtl")]
        public Output<int> IdentityTokenTtl { get; private set; } = null!;

        /// <summary>
        /// Number of max retries the client should use for recoverable errors. 
        /// The default `-1` falls back to the AWS SDK's default behavior.
        /// </summary>
        [Output("maxRetries")]
        public Output<int?> MaxRetries { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Role ARN to assume for plugin identity token federation. Requires Vault 1.17+.
        /// *Available only for Vault Enterprise*
        /// </summary>
        [Output("roleArn")]
        public Output<string?> RoleArn { get; private set; } = null!;

        /// <summary>
        /// The AWS secret key that Vault should use for the
        /// auth backend.
        /// </summary>
        [Output("secretKey")]
        public Output<string?> SecretKey { get; private set; } = null!;

        /// <summary>
        /// Override the URL Vault uses when making STS API
        /// calls.
        /// </summary>
        [Output("stsEndpoint")]
        public Output<string?> StsEndpoint { get; private set; } = null!;

        /// <summary>
        /// Override the default region when making STS API 
        /// calls. The `sts_endpoint` argument must be set when using `sts_region`.
        /// </summary>
        [Output("stsRegion")]
        public Output<string?> StsRegion { get; private set; } = null!;

        /// <summary>
        /// Available in Vault v1.15+. If set, 
        /// overrides both `sts_endpoint` and `sts_region` to instead use the region
        /// specified in the client request headers for IAM-based authentication.
        /// This can be useful when you have client requests coming from different
        /// regions and want flexibility in which regional STS API is used.
        /// </summary>
        [Output("useStsRegionFromClient")]
        public Output<bool> UseStsRegionFromClient { get; private set; } = null!;


        /// <summary>
        /// Create a AuthBackendClient resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackendClient(string name, AuthBackendClientArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:aws/authBackendClient:AuthBackendClient", name, args ?? new AuthBackendClientArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthBackendClient(string name, Input<string> id, AuthBackendClientState? state = null, CustomResourceOptions? options = null)
            : base("vault:aws/authBackendClient:AuthBackendClient", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "accessKey",
                    "secretKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AuthBackendClient resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthBackendClient Get(string name, Input<string> id, AuthBackendClientState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthBackendClient(name, id, state, options);
        }
    }

    public sealed class AuthBackendClientArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessKey")]
        private Input<string>? _accessKey;

        /// <summary>
        /// The AWS access key that Vault should use for the
        /// auth backend. Mutually exclusive with `identity_token_audience`.
        /// </summary>
        public Input<string>? AccessKey
        {
            get => _accessKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The path the AWS auth backend being configured was
        /// mounted at.  Defaults to `aws`.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// Override the URL Vault uses when making EC2 API
        /// calls.
        /// </summary>
        [Input("ec2Endpoint")]
        public Input<string>? Ec2Endpoint { get; set; }

        /// <summary>
        /// Override the URL Vault uses when making IAM API
        /// calls.
        /// </summary>
        [Input("iamEndpoint")]
        public Input<string>? IamEndpoint { get; set; }

        /// <summary>
        /// The value to require in the
        /// `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
        /// that are used in the IAM auth method.
        /// </summary>
        [Input("iamServerIdHeaderValue")]
        public Input<string>? IamServerIdHeaderValue { get; set; }

        /// <summary>
        /// The audience claim value. Mutually exclusive with `access_key`. 
        /// Requires Vault 1.17+. *Available only for Vault Enterprise*
        /// </summary>
        [Input("identityTokenAudience")]
        public Input<string>? IdentityTokenAudience { get; set; }

        /// <summary>
        /// The TTL of generated identity tokens in seconds. Requires Vault 1.17+.
        /// *Available only for Vault Enterprise*
        /// </summary>
        [Input("identityTokenTtl")]
        public Input<int>? IdentityTokenTtl { get; set; }

        /// <summary>
        /// Number of max retries the client should use for recoverable errors. 
        /// The default `-1` falls back to the AWS SDK's default behavior.
        /// </summary>
        [Input("maxRetries")]
        public Input<int>? MaxRetries { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Role ARN to assume for plugin identity token federation. Requires Vault 1.17+.
        /// *Available only for Vault Enterprise*
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("secretKey")]
        private Input<string>? _secretKey;

        /// <summary>
        /// The AWS secret key that Vault should use for the
        /// auth backend.
        /// </summary>
        public Input<string>? SecretKey
        {
            get => _secretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Override the URL Vault uses when making STS API
        /// calls.
        /// </summary>
        [Input("stsEndpoint")]
        public Input<string>? StsEndpoint { get; set; }

        /// <summary>
        /// Override the default region when making STS API 
        /// calls. The `sts_endpoint` argument must be set when using `sts_region`.
        /// </summary>
        [Input("stsRegion")]
        public Input<string>? StsRegion { get; set; }

        /// <summary>
        /// Available in Vault v1.15+. If set, 
        /// overrides both `sts_endpoint` and `sts_region` to instead use the region
        /// specified in the client request headers for IAM-based authentication.
        /// This can be useful when you have client requests coming from different
        /// regions and want flexibility in which regional STS API is used.
        /// </summary>
        [Input("useStsRegionFromClient")]
        public Input<bool>? UseStsRegionFromClient { get; set; }

        public AuthBackendClientArgs()
        {
        }
        public static new AuthBackendClientArgs Empty => new AuthBackendClientArgs();
    }

    public sealed class AuthBackendClientState : global::Pulumi.ResourceArgs
    {
        [Input("accessKey")]
        private Input<string>? _accessKey;

        /// <summary>
        /// The AWS access key that Vault should use for the
        /// auth backend. Mutually exclusive with `identity_token_audience`.
        /// </summary>
        public Input<string>? AccessKey
        {
            get => _accessKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The path the AWS auth backend being configured was
        /// mounted at.  Defaults to `aws`.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// Override the URL Vault uses when making EC2 API
        /// calls.
        /// </summary>
        [Input("ec2Endpoint")]
        public Input<string>? Ec2Endpoint { get; set; }

        /// <summary>
        /// Override the URL Vault uses when making IAM API
        /// calls.
        /// </summary>
        [Input("iamEndpoint")]
        public Input<string>? IamEndpoint { get; set; }

        /// <summary>
        /// The value to require in the
        /// `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
        /// that are used in the IAM auth method.
        /// </summary>
        [Input("iamServerIdHeaderValue")]
        public Input<string>? IamServerIdHeaderValue { get; set; }

        /// <summary>
        /// The audience claim value. Mutually exclusive with `access_key`. 
        /// Requires Vault 1.17+. *Available only for Vault Enterprise*
        /// </summary>
        [Input("identityTokenAudience")]
        public Input<string>? IdentityTokenAudience { get; set; }

        /// <summary>
        /// The TTL of generated identity tokens in seconds. Requires Vault 1.17+.
        /// *Available only for Vault Enterprise*
        /// </summary>
        [Input("identityTokenTtl")]
        public Input<int>? IdentityTokenTtl { get; set; }

        /// <summary>
        /// Number of max retries the client should use for recoverable errors. 
        /// The default `-1` falls back to the AWS SDK's default behavior.
        /// </summary>
        [Input("maxRetries")]
        public Input<int>? MaxRetries { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Role ARN to assume for plugin identity token federation. Requires Vault 1.17+.
        /// *Available only for Vault Enterprise*
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("secretKey")]
        private Input<string>? _secretKey;

        /// <summary>
        /// The AWS secret key that Vault should use for the
        /// auth backend.
        /// </summary>
        public Input<string>? SecretKey
        {
            get => _secretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Override the URL Vault uses when making STS API
        /// calls.
        /// </summary>
        [Input("stsEndpoint")]
        public Input<string>? StsEndpoint { get; set; }

        /// <summary>
        /// Override the default region when making STS API 
        /// calls. The `sts_endpoint` argument must be set when using `sts_region`.
        /// </summary>
        [Input("stsRegion")]
        public Input<string>? StsRegion { get; set; }

        /// <summary>
        /// Available in Vault v1.15+. If set, 
        /// overrides both `sts_endpoint` and `sts_region` to instead use the region
        /// specified in the client request headers for IAM-based authentication.
        /// This can be useful when you have client requests coming from different
        /// regions and want flexibility in which regional STS API is used.
        /// </summary>
        [Input("useStsRegionFromClient")]
        public Input<bool>? UseStsRegionFromClient { get; set; }

        public AuthBackendClientState()
        {
        }
        public static new AuthBackendClientState Empty => new AuthBackendClientState();
    }
}
