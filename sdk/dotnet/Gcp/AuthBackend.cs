// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
    /// 
    /// ## Example Usage
    /// 
    /// You can setup the GCP auth backend with Workload Identity Federation (WIF) for a secret-less configuration:
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var gcp = new Vault.Gcp.AuthBackend("gcp", new()
    ///     {
    ///         IdentityTokenKey = "example-key",
    ///         IdentityTokenTtl = 1800,
    ///         IdentityTokenAudience = "&lt;TOKEN_AUDIENCE&gt;",
    ///         ServiceAccountEmail = "&lt;SERVICE_ACCOUNT_EMAIL&gt;",
    ///         RotationSchedule = "0 * * * SAT",
    ///         RotationWindow = 3600,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// GCP authentication backends can be imported using the backend name, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:gcp/authBackend:AuthBackend gcp gcp
    /// ```
    /// </summary>
    [VaultResourceType("vault:gcp/authBackend:AuthBackend")]
    public partial class AuthBackend : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
        /// </summary>
        [Output("accessor")]
        public Output<string> Accessor { get; private set; } = null!;

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

        /// <summary>
        /// Specifies overrides to
        /// [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
        /// used when making API requests. This allows specific requests made during authentication
        /// to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
        /// environments. Requires Vault 1.11+.
        /// 
        /// Overrides are set at the subdomain level using the following keys:
        /// </summary>
        [Output("customEndpoint")]
        public Output<Outputs.AuthBackendCustomEndpoint?> CustomEndpoint { get; private set; } = null!;

        /// <summary>
        /// A description of the auth method.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
        /// </summary>
        [Output("disableAutomatedRotation")]
        public Output<bool?> DisableAutomatedRotation { get; private set; } = null!;

        /// <summary>
        /// If set, opts out of mount migration on path updates.
        /// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        /// </summary>
        [Output("disableRemount")]
        public Output<bool?> DisableRemount { get; private set; } = null!;

        /// <summary>
        /// The audience claim value for plugin identity
        /// tokens. Must match an allowed audience configured for the target [Workload Identity Pool](https://cloud.google.com/iam/docs/workload-identity-federation-with-other-providers#prepare).
        /// Mutually exclusive with `credentials`.  Requires Vault 1.17+. *Available only for Vault Enterprise*.
        /// </summary>
        [Output("identityTokenAudience")]
        public Output<string?> IdentityTokenAudience { get; private set; } = null!;

        /// <summary>
        /// The key to use for signing plugin identity
        /// tokens. Requires Vault 1.17+. *Available only for Vault Enterprise*.
        /// </summary>
        [Output("identityTokenKey")]
        public Output<string?> IdentityTokenKey { get; private set; } = null!;

        /// <summary>
        /// The TTL of generated tokens.
        /// </summary>
        [Output("identityTokenTtl")]
        public Output<int?> IdentityTokenTtl { get; private set; } = null!;

        /// <summary>
        /// Specifies if the auth method is local only.
        /// </summary>
        [Output("local")]
        public Output<bool?> Local { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// The path to mount the auth method — this defaults to 'gcp'.
        /// </summary>
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
        /// The amount of time in seconds Vault should wait before rotating the root credential.
        /// A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
        /// </summary>
        [Output("rotationPeriod")]
        public Output<int?> RotationPeriod { get; private set; } = null!;

        /// <summary>
        /// The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
        /// defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
        /// </summary>
        [Output("rotationSchedule")]
        public Output<string?> RotationSchedule { get; private set; } = null!;

        /// <summary>
        /// The maximum amount of time in seconds allowed to complete
        /// a rotation when a scheduled token rotation occurs. The default rotation window is
        /// unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
        /// </summary>
        [Output("rotationWindow")]
        public Output<int?> RotationWindow { get; private set; } = null!;

        /// <summary>
        /// Service Account to impersonate for plugin workload identity federation.
        /// Required with `identity_token_audience`. Requires Vault 1.17+. *Available only for Vault Enterprise*.
        /// </summary>
        [Output("serviceAccountEmail")]
        public Output<string?> ServiceAccountEmail { get; private set; } = null!;

        /// <summary>
        /// Extra configuration block. Structure is documented below.
        /// 
        /// The `tune` block is used to tune the auth backend:
        /// </summary>
        [Output("tune")]
        public Output<Outputs.AuthBackendTune> Tune { get; private set; } = null!;


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
                AdditionalSecretOutputs =
                {
                    "credentials",
                },
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

    public sealed class AuthBackendArgs : global::Pulumi.ResourceArgs
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

        [Input("credentials")]
        private Input<string>? _credentials;

        /// <summary>
        /// A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
        /// </summary>
        public Input<string>? Credentials
        {
            get => _credentials;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _credentials = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specifies overrides to
        /// [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
        /// used when making API requests. This allows specific requests made during authentication
        /// to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
        /// environments. Requires Vault 1.11+.
        /// 
        /// Overrides are set at the subdomain level using the following keys:
        /// </summary>
        [Input("customEndpoint")]
        public Input<Inputs.AuthBackendCustomEndpointArgs>? CustomEndpoint { get; set; }

        /// <summary>
        /// A description of the auth method.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
        /// </summary>
        [Input("disableAutomatedRotation")]
        public Input<bool>? DisableAutomatedRotation { get; set; }

        /// <summary>
        /// If set, opts out of mount migration on path updates.
        /// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        /// </summary>
        [Input("disableRemount")]
        public Input<bool>? DisableRemount { get; set; }

        /// <summary>
        /// The audience claim value for plugin identity
        /// tokens. Must match an allowed audience configured for the target [Workload Identity Pool](https://cloud.google.com/iam/docs/workload-identity-federation-with-other-providers#prepare).
        /// Mutually exclusive with `credentials`.  Requires Vault 1.17+. *Available only for Vault Enterprise*.
        /// </summary>
        [Input("identityTokenAudience")]
        public Input<string>? IdentityTokenAudience { get; set; }

        /// <summary>
        /// The key to use for signing plugin identity
        /// tokens. Requires Vault 1.17+. *Available only for Vault Enterprise*.
        /// </summary>
        [Input("identityTokenKey")]
        public Input<string>? IdentityTokenKey { get; set; }

        /// <summary>
        /// The TTL of generated tokens.
        /// </summary>
        [Input("identityTokenTtl")]
        public Input<int>? IdentityTokenTtl { get; set; }

        /// <summary>
        /// Specifies if the auth method is local only.
        /// </summary>
        [Input("local")]
        public Input<bool>? Local { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The path to mount the auth method — this defaults to 'gcp'.
        /// </summary>
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

        /// <summary>
        /// The amount of time in seconds Vault should wait before rotating the root credential.
        /// A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
        /// </summary>
        [Input("rotationPeriod")]
        public Input<int>? RotationPeriod { get; set; }

        /// <summary>
        /// The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
        /// defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
        /// </summary>
        [Input("rotationSchedule")]
        public Input<string>? RotationSchedule { get; set; }

        /// <summary>
        /// The maximum amount of time in seconds allowed to complete
        /// a rotation when a scheduled token rotation occurs. The default rotation window is
        /// unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
        /// </summary>
        [Input("rotationWindow")]
        public Input<int>? RotationWindow { get; set; }

        /// <summary>
        /// Service Account to impersonate for plugin workload identity federation.
        /// Required with `identity_token_audience`. Requires Vault 1.17+. *Available only for Vault Enterprise*.
        /// </summary>
        [Input("serviceAccountEmail")]
        public Input<string>? ServiceAccountEmail { get; set; }

        /// <summary>
        /// Extra configuration block. Structure is documented below.
        /// 
        /// The `tune` block is used to tune the auth backend:
        /// </summary>
        [Input("tune")]
        public Input<Inputs.AuthBackendTuneArgs>? Tune { get; set; }

        public AuthBackendArgs()
        {
        }
        public static new AuthBackendArgs Empty => new AuthBackendArgs();
    }

    public sealed class AuthBackendState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
        /// </summary>
        [Input("accessor")]
        public Input<string>? Accessor { get; set; }

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

        [Input("credentials")]
        private Input<string>? _credentials;

        /// <summary>
        /// A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
        /// </summary>
        public Input<string>? Credentials
        {
            get => _credentials;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _credentials = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specifies overrides to
        /// [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
        /// used when making API requests. This allows specific requests made during authentication
        /// to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
        /// environments. Requires Vault 1.11+.
        /// 
        /// Overrides are set at the subdomain level using the following keys:
        /// </summary>
        [Input("customEndpoint")]
        public Input<Inputs.AuthBackendCustomEndpointGetArgs>? CustomEndpoint { get; set; }

        /// <summary>
        /// A description of the auth method.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Cancels all upcoming rotations of the root credential until unset. Requires Vault Enterprise 1.19+.
        /// </summary>
        [Input("disableAutomatedRotation")]
        public Input<bool>? DisableAutomatedRotation { get; set; }

        /// <summary>
        /// If set, opts out of mount migration on path updates.
        /// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        /// </summary>
        [Input("disableRemount")]
        public Input<bool>? DisableRemount { get; set; }

        /// <summary>
        /// The audience claim value for plugin identity
        /// tokens. Must match an allowed audience configured for the target [Workload Identity Pool](https://cloud.google.com/iam/docs/workload-identity-federation-with-other-providers#prepare).
        /// Mutually exclusive with `credentials`.  Requires Vault 1.17+. *Available only for Vault Enterprise*.
        /// </summary>
        [Input("identityTokenAudience")]
        public Input<string>? IdentityTokenAudience { get; set; }

        /// <summary>
        /// The key to use for signing plugin identity
        /// tokens. Requires Vault 1.17+. *Available only for Vault Enterprise*.
        /// </summary>
        [Input("identityTokenKey")]
        public Input<string>? IdentityTokenKey { get; set; }

        /// <summary>
        /// The TTL of generated tokens.
        /// </summary>
        [Input("identityTokenTtl")]
        public Input<int>? IdentityTokenTtl { get; set; }

        /// <summary>
        /// Specifies if the auth method is local only.
        /// </summary>
        [Input("local")]
        public Input<bool>? Local { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The path to mount the auth method — this defaults to 'gcp'.
        /// </summary>
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

        /// <summary>
        /// The amount of time in seconds Vault should wait before rotating the root credential.
        /// A zero value tells Vault not to rotate the root credential. The minimum rotation period is 10 seconds. Requires Vault Enterprise 1.19+.
        /// </summary>
        [Input("rotationPeriod")]
        public Input<int>? RotationPeriod { get; set; }

        /// <summary>
        /// The schedule, in [cron-style time format](https://en.wikipedia.org/wiki/Cron),
        /// defining the schedule on which Vault should rotate the root token. Requires Vault Enterprise 1.19+.
        /// </summary>
        [Input("rotationSchedule")]
        public Input<string>? RotationSchedule { get; set; }

        /// <summary>
        /// The maximum amount of time in seconds allowed to complete
        /// a rotation when a scheduled token rotation occurs. The default rotation window is
        /// unbound and the minimum allowable window is `3600`. Requires Vault Enterprise 1.19+.
        /// </summary>
        [Input("rotationWindow")]
        public Input<int>? RotationWindow { get; set; }

        /// <summary>
        /// Service Account to impersonate for plugin workload identity federation.
        /// Required with `identity_token_audience`. Requires Vault 1.17+. *Available only for Vault Enterprise*.
        /// </summary>
        [Input("serviceAccountEmail")]
        public Input<string>? ServiceAccountEmail { get; set; }

        /// <summary>
        /// Extra configuration block. Structure is documented below.
        /// 
        /// The `tune` block is used to tune the auth backend:
        /// </summary>
        [Input("tune")]
        public Input<Inputs.AuthBackendTuneGetArgs>? Tune { get; set; }

        public AuthBackendState()
        {
        }
        public static new AuthBackendState Empty => new AuthBackendState();
    }
}
