// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Jwt
{
    /// <summary>
    /// Provides a resource for managing an
    /// [JWT auth backend within Vault](https://www.vaultproject.io/docs/auth/jwt.html).
    /// 
    /// ## Example Usage
    /// 
    /// Manage JWT auth backend:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Vault.Jwt.AuthBackend("example", new Vault.Jwt.AuthBackendArgs
    ///         {
    ///             BoundIssuer = "https://myco.auth0.com/",
    ///             Description = "Demonstration of the Terraform JWT auth backend",
    ///             OidcDiscoveryUrl = "https://myco.auth0.com/",
    ///             Path = "jwt",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// Manage OIDC auth backend:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Vault.Jwt.AuthBackend("example", new Vault.Jwt.AuthBackendArgs
    ///         {
    ///             BoundIssuer = "https://myco.auth0.com/",
    ///             Description = "Demonstration of the Terraform JWT auth backend",
    ///             OidcClientId = "1234567890",
    ///             OidcClientSecret = "secret123456",
    ///             OidcDiscoveryUrl = "https://myco.auth0.com/",
    ///             Path = "oidc",
    ///             Tune = new Vault.Jwt.Inputs.AuthBackendTuneArgs
    ///             {
    ///                 ListingVisibility = "unauth",
    ///             },
    ///             Type = "oidc",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class AuthBackend : Pulumi.CustomResource
    {
        /// <summary>
        /// The accessor of the JWT auth backend
        /// </summary>
        [Output("accessor")]
        public Output<string> Accessor { get; private set; } = null!;

        /// <summary>
        /// The value against which to match the iss claim in a JWT
        /// </summary>
        [Output("boundIssuer")]
        public Output<string?> BoundIssuer { get; private set; } = null!;

        /// <summary>
        /// The default role to use if none is provided during login
        /// </summary>
        [Output("defaultRole")]
        public Output<string?> DefaultRole { get; private set; } = null!;

        /// <summary>
        /// The description of the auth backend
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
        /// </summary>
        [Output("jwksCaPem")]
        public Output<string?> JwksCaPem { get; private set; } = null!;

        /// <summary>
        /// JWKS URL to use to authenticate signatures. Cannot be used with "oidc_discovery_url" or "jwt_validation_pubkeys".
        /// </summary>
        [Output("jwksUrl")]
        public Output<string?> JwksUrl { get; private set; } = null!;

        /// <summary>
        /// A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ
        /// </summary>
        [Output("jwtSupportedAlgs")]
        public Output<ImmutableArray<string>> JwtSupportedAlgs { get; private set; } = null!;

        /// <summary>
        /// A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with `oidc_discovery_url`
        /// </summary>
        [Output("jwtValidationPubkeys")]
        public Output<ImmutableArray<string>> JwtValidationPubkeys { get; private set; } = null!;

        /// <summary>
        /// Client ID used for OIDC backends
        /// </summary>
        [Output("oidcClientId")]
        public Output<string?> OidcClientId { get; private set; } = null!;

        /// <summary>
        /// Client Secret used for OIDC backends
        /// </summary>
        [Output("oidcClientSecret")]
        public Output<string?> OidcClientSecret { get; private set; } = null!;

        /// <summary>
        /// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
        /// </summary>
        [Output("oidcDiscoveryCaPem")]
        public Output<string?> OidcDiscoveryCaPem { get; private set; } = null!;

        /// <summary>
        /// The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwt_validation_pubkeys`
        /// </summary>
        [Output("oidcDiscoveryUrl")]
        public Output<string?> OidcDiscoveryUrl { get; private set; } = null!;

        /// <summary>
        /// Path to mount the JWT/OIDC auth backend
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        [Output("tune")]
        public Output<Outputs.AuthBackendTune> Tune { get; private set; } = null!;

        /// <summary>
        /// Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AuthBackend resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackend(string name, AuthBackendArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:jwt/authBackend:AuthBackend", name, args ?? new AuthBackendArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthBackend(string name, Input<string> id, AuthBackendState? state = null, CustomResourceOptions? options = null)
            : base("vault:jwt/authBackend:AuthBackend", name, state, MakeResourceOptions(options, id))
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
        /// The value against which to match the iss claim in a JWT
        /// </summary>
        [Input("boundIssuer")]
        public Input<string>? BoundIssuer { get; set; }

        /// <summary>
        /// The default role to use if none is provided during login
        /// </summary>
        [Input("defaultRole")]
        public Input<string>? DefaultRole { get; set; }

        /// <summary>
        /// The description of the auth backend
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
        /// </summary>
        [Input("jwksCaPem")]
        public Input<string>? JwksCaPem { get; set; }

        /// <summary>
        /// JWKS URL to use to authenticate signatures. Cannot be used with "oidc_discovery_url" or "jwt_validation_pubkeys".
        /// </summary>
        [Input("jwksUrl")]
        public Input<string>? JwksUrl { get; set; }

        [Input("jwtSupportedAlgs")]
        private InputList<string>? _jwtSupportedAlgs;

        /// <summary>
        /// A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ
        /// </summary>
        public InputList<string> JwtSupportedAlgs
        {
            get => _jwtSupportedAlgs ?? (_jwtSupportedAlgs = new InputList<string>());
            set => _jwtSupportedAlgs = value;
        }

        [Input("jwtValidationPubkeys")]
        private InputList<string>? _jwtValidationPubkeys;

        /// <summary>
        /// A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with `oidc_discovery_url`
        /// </summary>
        public InputList<string> JwtValidationPubkeys
        {
            get => _jwtValidationPubkeys ?? (_jwtValidationPubkeys = new InputList<string>());
            set => _jwtValidationPubkeys = value;
        }

        /// <summary>
        /// Client ID used for OIDC backends
        /// </summary>
        [Input("oidcClientId")]
        public Input<string>? OidcClientId { get; set; }

        /// <summary>
        /// Client Secret used for OIDC backends
        /// </summary>
        [Input("oidcClientSecret")]
        public Input<string>? OidcClientSecret { get; set; }

        /// <summary>
        /// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
        /// </summary>
        [Input("oidcDiscoveryCaPem")]
        public Input<string>? OidcDiscoveryCaPem { get; set; }

        /// <summary>
        /// The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwt_validation_pubkeys`
        /// </summary>
        [Input("oidcDiscoveryUrl")]
        public Input<string>? OidcDiscoveryUrl { get; set; }

        /// <summary>
        /// Path to mount the JWT/OIDC auth backend
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("tune")]
        public Input<Inputs.AuthBackendTuneArgs>? Tune { get; set; }

        /// <summary>
        /// Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AuthBackendArgs()
        {
        }
    }

    public sealed class AuthBackendState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The accessor of the JWT auth backend
        /// </summary>
        [Input("accessor")]
        public Input<string>? Accessor { get; set; }

        /// <summary>
        /// The value against which to match the iss claim in a JWT
        /// </summary>
        [Input("boundIssuer")]
        public Input<string>? BoundIssuer { get; set; }

        /// <summary>
        /// The default role to use if none is provided during login
        /// </summary>
        [Input("defaultRole")]
        public Input<string>? DefaultRole { get; set; }

        /// <summary>
        /// The description of the auth backend
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
        /// </summary>
        [Input("jwksCaPem")]
        public Input<string>? JwksCaPem { get; set; }

        /// <summary>
        /// JWKS URL to use to authenticate signatures. Cannot be used with "oidc_discovery_url" or "jwt_validation_pubkeys".
        /// </summary>
        [Input("jwksUrl")]
        public Input<string>? JwksUrl { get; set; }

        [Input("jwtSupportedAlgs")]
        private InputList<string>? _jwtSupportedAlgs;

        /// <summary>
        /// A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ
        /// </summary>
        public InputList<string> JwtSupportedAlgs
        {
            get => _jwtSupportedAlgs ?? (_jwtSupportedAlgs = new InputList<string>());
            set => _jwtSupportedAlgs = value;
        }

        [Input("jwtValidationPubkeys")]
        private InputList<string>? _jwtValidationPubkeys;

        /// <summary>
        /// A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with `oidc_discovery_url`
        /// </summary>
        public InputList<string> JwtValidationPubkeys
        {
            get => _jwtValidationPubkeys ?? (_jwtValidationPubkeys = new InputList<string>());
            set => _jwtValidationPubkeys = value;
        }

        /// <summary>
        /// Client ID used for OIDC backends
        /// </summary>
        [Input("oidcClientId")]
        public Input<string>? OidcClientId { get; set; }

        /// <summary>
        /// Client Secret used for OIDC backends
        /// </summary>
        [Input("oidcClientSecret")]
        public Input<string>? OidcClientSecret { get; set; }

        /// <summary>
        /// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
        /// </summary>
        [Input("oidcDiscoveryCaPem")]
        public Input<string>? OidcDiscoveryCaPem { get; set; }

        /// <summary>
        /// The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwt_validation_pubkeys`
        /// </summary>
        [Input("oidcDiscoveryUrl")]
        public Input<string>? OidcDiscoveryUrl { get; set; }

        /// <summary>
        /// Path to mount the JWT/OIDC auth backend
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("tune")]
        public Input<Inputs.AuthBackendTuneGetArgs>? Tune { get; set; }

        /// <summary>
        /// Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AuthBackendState()
        {
        }
    }
}
