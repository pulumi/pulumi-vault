// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Consul
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ### Creating a standard backend resource:
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test = new Vault.Consul.SecretBackend("test", new()
    ///     {
    ///         Path = "consul",
    ///         Description = "Manages the Consul backend",
    ///         Address = "127.0.0.1:8500",
    ///         Token = "4240861b-ce3d-8530-115a-521ff070dd29",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Creating a backend resource to bootstrap a new Consul instance:
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test = new Vault.Consul.SecretBackend("test", new()
    ///     {
    ///         Path = "consul",
    ///         Description = "Bootstrap the Consul backend",
    ///         Address = "127.0.0.1:8500",
    ///         Bootstrap = true,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Consul secret backends can be imported using the `path`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:consul/secretBackend:SecretBackend example consul
    /// ```
    /// </summary>
    [VaultResourceType("vault:consul/secretBackend:SecretBackend")]
    public partial class SecretBackend : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// Denotes a backend resource that is used to bootstrap the Consul ACL system. Only one resource may be used to bootstrap.
        /// </summary>
        [Output("bootstrap")]
        public Output<bool?> Bootstrap { get; private set; } = null!;

        /// <summary>
        /// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
        /// </summary>
        [Output("caCert")]
        public Output<string?> CaCert { get; private set; } = null!;

        /// <summary>
        /// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if
        /// this is set you need to also set client_key.
        /// </summary>
        [Output("clientCert")]
        public Output<string?> ClientCert { get; private set; } = null!;

        /// <summary>
        /// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set
        /// you need to also set client_cert.
        /// </summary>
        [Output("clientKey")]
        public Output<string?> ClientKey { get; private set; } = null!;

        /// <summary>
        /// The default TTL for credentials issued by this backend.
        /// </summary>
        [Output("defaultLeaseTtlSeconds")]
        public Output<int?> DefaultLeaseTtlSeconds { get; private set; } = null!;

        /// <summary>
        /// A human-friendly description for this backend.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// If set, opts out of mount migration on path updates.
        /// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        /// </summary>
        [Output("disableRemount")]
        public Output<bool?> DisableRemount { get; private set; } = null!;

        /// <summary>
        /// Specifies if the secret backend is local only.
        /// </summary>
        [Output("local")]
        public Output<bool?> Local { get; private set; } = null!;

        /// <summary>
        /// The maximum TTL that can be requested
        /// for credentials issued by this backend.
        /// </summary>
        [Output("maxLeaseTtlSeconds")]
        public Output<int?> MaxLeaseTtlSeconds { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults
        /// to `consul`.
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        /// <summary>
        /// Specifies the URL scheme to use. Defaults to `http`.
        /// </summary>
        [Output("scheme")]
        public Output<string?> Scheme { get; private set; } = null!;

        /// <summary>
        /// Specifies the Consul token to use when managing or issuing new tokens.
        /// </summary>
        [Output("token")]
        public Output<string?> Token { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackend resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackend(string name, SecretBackendArgs args, CustomResourceOptions? options = null)
            : base("vault:consul/secretBackend:SecretBackend", name, args ?? new SecretBackendArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackend(string name, Input<string> id, SecretBackendState? state = null, CustomResourceOptions? options = null)
            : base("vault:consul/secretBackend:SecretBackend", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "clientCert",
                    "clientKey",
                    "token",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecretBackend resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackend Get(string name, Input<string> id, SecretBackendState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackend(name, id, state, options);
        }
    }

    public sealed class SecretBackendArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
        /// </summary>
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        /// <summary>
        /// Denotes a backend resource that is used to bootstrap the Consul ACL system. Only one resource may be used to bootstrap.
        /// </summary>
        [Input("bootstrap")]
        public Input<bool>? Bootstrap { get; set; }

        /// <summary>
        /// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
        /// </summary>
        [Input("caCert")]
        public Input<string>? CaCert { get; set; }

        [Input("clientCert")]
        private Input<string>? _clientCert;

        /// <summary>
        /// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if
        /// this is set you need to also set client_key.
        /// </summary>
        public Input<string>? ClientCert
        {
            get => _clientCert;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientCert = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("clientKey")]
        private Input<string>? _clientKey;

        /// <summary>
        /// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set
        /// you need to also set client_cert.
        /// </summary>
        public Input<string>? ClientKey
        {
            get => _clientKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The default TTL for credentials issued by this backend.
        /// </summary>
        [Input("defaultLeaseTtlSeconds")]
        public Input<int>? DefaultLeaseTtlSeconds { get; set; }

        /// <summary>
        /// A human-friendly description for this backend.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If set, opts out of mount migration on path updates.
        /// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        /// </summary>
        [Input("disableRemount")]
        public Input<bool>? DisableRemount { get; set; }

        /// <summary>
        /// Specifies if the secret backend is local only.
        /// </summary>
        [Input("local")]
        public Input<bool>? Local { get; set; }

        /// <summary>
        /// The maximum TTL that can be requested
        /// for credentials issued by this backend.
        /// </summary>
        [Input("maxLeaseTtlSeconds")]
        public Input<int>? MaxLeaseTtlSeconds { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults
        /// to `consul`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Specifies the URL scheme to use. Defaults to `http`.
        /// </summary>
        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        [Input("token")]
        private Input<string>? _token;

        /// <summary>
        /// Specifies the Consul token to use when managing or issuing new tokens.
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

        public SecretBackendArgs()
        {
        }
        public static new SecretBackendArgs Empty => new SecretBackendArgs();
    }

    public sealed class SecretBackendState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Denotes a backend resource that is used to bootstrap the Consul ACL system. Only one resource may be used to bootstrap.
        /// </summary>
        [Input("bootstrap")]
        public Input<bool>? Bootstrap { get; set; }

        /// <summary>
        /// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
        /// </summary>
        [Input("caCert")]
        public Input<string>? CaCert { get; set; }

        [Input("clientCert")]
        private Input<string>? _clientCert;

        /// <summary>
        /// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if
        /// this is set you need to also set client_key.
        /// </summary>
        public Input<string>? ClientCert
        {
            get => _clientCert;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientCert = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("clientKey")]
        private Input<string>? _clientKey;

        /// <summary>
        /// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set
        /// you need to also set client_cert.
        /// </summary>
        public Input<string>? ClientKey
        {
            get => _clientKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The default TTL for credentials issued by this backend.
        /// </summary>
        [Input("defaultLeaseTtlSeconds")]
        public Input<int>? DefaultLeaseTtlSeconds { get; set; }

        /// <summary>
        /// A human-friendly description for this backend.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If set, opts out of mount migration on path updates.
        /// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        /// </summary>
        [Input("disableRemount")]
        public Input<bool>? DisableRemount { get; set; }

        /// <summary>
        /// Specifies if the secret backend is local only.
        /// </summary>
        [Input("local")]
        public Input<bool>? Local { get; set; }

        /// <summary>
        /// The maximum TTL that can be requested
        /// for credentials issued by this backend.
        /// </summary>
        [Input("maxLeaseTtlSeconds")]
        public Input<int>? MaxLeaseTtlSeconds { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults
        /// to `consul`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Specifies the URL scheme to use. Defaults to `http`.
        /// </summary>
        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        [Input("token")]
        private Input<string>? _token;

        /// <summary>
        /// Specifies the Consul token to use when managing or issuing new tokens.
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

        public SecretBackendState()
        {
        }
        public static new SecretBackendState Empty => new SecretBackendState();
    }
}
