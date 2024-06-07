// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Kmip
{
    /// <summary>
    /// Manages KMIP Secret backends in a Vault server. This feature requires
    /// Vault Enterprise. See the [Vault documentation](https://www.vaultproject.io/docs/secrets/kmip)
    /// for more information.
    /// 
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
    ///     var @default = new Vault.Kmip.SecretBackend("default", new()
    ///     {
    ///         Path = "kmip",
    ///         Description = "Vault KMIP backend",
    ///         ListenAddrs = new[]
    ///         {
    ///             "127.0.0.1:5696",
    ///             "127.0.0.1:8080",
    ///         },
    ///         TlsCaKeyType = "rsa",
    ///         TlsCaKeyBits = 4096,
    ///         DefaultTlsClientKeyType = "rsa",
    ///         DefaultTlsClientKeyBits = 4096,
    ///         DefaultTlsClientTtl = 86400,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// KMIP Secret backend can be imported using the `path`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:kmip/secretBackend:SecretBackend default kmip
    /// ```
    /// </summary>
    [VaultResourceType("vault:kmip/secretBackend:SecretBackend")]
    public partial class SecretBackend : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Client certificate key bits, valid values depend on key type.
        /// </summary>
        [Output("defaultTlsClientKeyBits")]
        public Output<int> DefaultTlsClientKeyBits { get; private set; } = null!;

        /// <summary>
        /// Client certificate key type, `rsa` or `ec`.
        /// </summary>
        [Output("defaultTlsClientKeyType")]
        public Output<string> DefaultTlsClientKeyType { get; private set; } = null!;

        /// <summary>
        /// Client certificate TTL in seconds
        /// </summary>
        [Output("defaultTlsClientTtl")]
        public Output<int> DefaultTlsClientTtl { get; private set; } = null!;

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
        /// Addresses the KMIP server should listen on (`host:port`).
        /// </summary>
        [Output("listenAddrs")]
        public Output<ImmutableArray<string>> ListenAddrs { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `kmip`.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// Hostnames to include in the server's TLS certificate as SAN DNS names. The first will be used as the common name (CN).
        /// </summary>
        [Output("serverHostnames")]
        public Output<ImmutableArray<string>> ServerHostnames { get; private set; } = null!;

        /// <summary>
        /// IPs to include in the server's TLS certificate as SAN IP addresses.
        /// </summary>
        [Output("serverIps")]
        public Output<ImmutableArray<string>> ServerIps { get; private set; } = null!;

        /// <summary>
        /// CA key bits, valid values depend on key type.
        /// </summary>
        [Output("tlsCaKeyBits")]
        public Output<int> TlsCaKeyBits { get; private set; } = null!;

        /// <summary>
        /// CA key type, rsa or ec.
        /// </summary>
        [Output("tlsCaKeyType")]
        public Output<string> TlsCaKeyType { get; private set; } = null!;

        /// <summary>
        /// Minimum TLS version to accept.
        /// </summary>
        [Output("tlsMinVersion")]
        public Output<string> TlsMinVersion { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackend resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackend(string name, SecretBackendArgs args, CustomResourceOptions? options = null)
            : base("vault:kmip/secretBackend:SecretBackend", name, args ?? new SecretBackendArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackend(string name, Input<string> id, SecretBackendState? state = null, CustomResourceOptions? options = null)
            : base("vault:kmip/secretBackend:SecretBackend", name, state, MakeResourceOptions(options, id))
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
        /// Client certificate key bits, valid values depend on key type.
        /// </summary>
        [Input("defaultTlsClientKeyBits")]
        public Input<int>? DefaultTlsClientKeyBits { get; set; }

        /// <summary>
        /// Client certificate key type, `rsa` or `ec`.
        /// </summary>
        [Input("defaultTlsClientKeyType")]
        public Input<string>? DefaultTlsClientKeyType { get; set; }

        /// <summary>
        /// Client certificate TTL in seconds
        /// </summary>
        [Input("defaultTlsClientTtl")]
        public Input<int>? DefaultTlsClientTtl { get; set; }

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

        [Input("listenAddrs")]
        private InputList<string>? _listenAddrs;

        /// <summary>
        /// Addresses the KMIP server should listen on (`host:port`).
        /// </summary>
        public InputList<string> ListenAddrs
        {
            get => _listenAddrs ?? (_listenAddrs = new InputList<string>());
            set => _listenAddrs = value;
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
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `kmip`.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        [Input("serverHostnames")]
        private InputList<string>? _serverHostnames;

        /// <summary>
        /// Hostnames to include in the server's TLS certificate as SAN DNS names. The first will be used as the common name (CN).
        /// </summary>
        public InputList<string> ServerHostnames
        {
            get => _serverHostnames ?? (_serverHostnames = new InputList<string>());
            set => _serverHostnames = value;
        }

        [Input("serverIps")]
        private InputList<string>? _serverIps;

        /// <summary>
        /// IPs to include in the server's TLS certificate as SAN IP addresses.
        /// </summary>
        public InputList<string> ServerIps
        {
            get => _serverIps ?? (_serverIps = new InputList<string>());
            set => _serverIps = value;
        }

        /// <summary>
        /// CA key bits, valid values depend on key type.
        /// </summary>
        [Input("tlsCaKeyBits")]
        public Input<int>? TlsCaKeyBits { get; set; }

        /// <summary>
        /// CA key type, rsa or ec.
        /// </summary>
        [Input("tlsCaKeyType")]
        public Input<string>? TlsCaKeyType { get; set; }

        /// <summary>
        /// Minimum TLS version to accept.
        /// </summary>
        [Input("tlsMinVersion")]
        public Input<string>? TlsMinVersion { get; set; }

        public SecretBackendArgs()
        {
        }
        public static new SecretBackendArgs Empty => new SecretBackendArgs();
    }

    public sealed class SecretBackendState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Client certificate key bits, valid values depend on key type.
        /// </summary>
        [Input("defaultTlsClientKeyBits")]
        public Input<int>? DefaultTlsClientKeyBits { get; set; }

        /// <summary>
        /// Client certificate key type, `rsa` or `ec`.
        /// </summary>
        [Input("defaultTlsClientKeyType")]
        public Input<string>? DefaultTlsClientKeyType { get; set; }

        /// <summary>
        /// Client certificate TTL in seconds
        /// </summary>
        [Input("defaultTlsClientTtl")]
        public Input<int>? DefaultTlsClientTtl { get; set; }

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

        [Input("listenAddrs")]
        private InputList<string>? _listenAddrs;

        /// <summary>
        /// Addresses the KMIP server should listen on (`host:port`).
        /// </summary>
        public InputList<string> ListenAddrs
        {
            get => _listenAddrs ?? (_listenAddrs = new InputList<string>());
            set => _listenAddrs = value;
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
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `kmip`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("serverHostnames")]
        private InputList<string>? _serverHostnames;

        /// <summary>
        /// Hostnames to include in the server's TLS certificate as SAN DNS names. The first will be used as the common name (CN).
        /// </summary>
        public InputList<string> ServerHostnames
        {
            get => _serverHostnames ?? (_serverHostnames = new InputList<string>());
            set => _serverHostnames = value;
        }

        [Input("serverIps")]
        private InputList<string>? _serverIps;

        /// <summary>
        /// IPs to include in the server's TLS certificate as SAN IP addresses.
        /// </summary>
        public InputList<string> ServerIps
        {
            get => _serverIps ?? (_serverIps = new InputList<string>());
            set => _serverIps = value;
        }

        /// <summary>
        /// CA key bits, valid values depend on key type.
        /// </summary>
        [Input("tlsCaKeyBits")]
        public Input<int>? TlsCaKeyBits { get; set; }

        /// <summary>
        /// CA key type, rsa or ec.
        /// </summary>
        [Input("tlsCaKeyType")]
        public Input<string>? TlsCaKeyType { get; set; }

        /// <summary>
        /// Minimum TLS version to accept.
        /// </summary>
        [Input("tlsMinVersion")]
        public Input<string>? TlsMinVersion { get; set; }

        public SecretBackendState()
        {
        }
        public static new SecretBackendState Empty => new SecretBackendState();
    }
}
