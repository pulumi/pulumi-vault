// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Vault
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("vault");

        private static readonly __Value<string?> _addAddressToEnv = new __Value<string?>(() => __config.Get("addAddressToEnv"));
        /// <summary>
        /// If true, adds the value of the `address` argument to the Terraform process environment.
        /// </summary>
        public static string? AddAddressToEnv
        {
            get => _addAddressToEnv.Get();
            set => _addAddressToEnv.Set(value);
        }

        private static readonly __Value<string?> _address = new __Value<string?>(() => __config.Get("address"));
        /// <summary>
        /// URL of the root of the target Vault server.
        /// </summary>
        public static string? Address
        {
            get => _address.Get();
            set => _address.Set(value);
        }

        private static readonly __Value<Pulumi.Vault.Config.Types.AuthLogin?> _authLogin = new __Value<Pulumi.Vault.Config.Types.AuthLogin?>(() => __config.GetObject<Pulumi.Vault.Config.Types.AuthLogin>("authLogin"));
        /// <summary>
        /// Login to vault with an existing auth method using auth/&lt;mount&gt;/login
        /// </summary>
        public static Pulumi.Vault.Config.Types.AuthLogin? AuthLogin
        {
            get => _authLogin.Get();
            set => _authLogin.Set(value);
        }

        private static readonly __Value<Pulumi.Vault.Config.Types.AuthLoginAws?> _authLoginAws = new __Value<Pulumi.Vault.Config.Types.AuthLoginAws?>(() => __config.GetObject<Pulumi.Vault.Config.Types.AuthLoginAws>("authLoginAws"));
        /// <summary>
        /// Login to vault using the AWS method
        /// </summary>
        public static Pulumi.Vault.Config.Types.AuthLoginAws? AuthLoginAws
        {
            get => _authLoginAws.Get();
            set => _authLoginAws.Set(value);
        }

        private static readonly __Value<Pulumi.Vault.Config.Types.AuthLoginAzure?> _authLoginAzure = new __Value<Pulumi.Vault.Config.Types.AuthLoginAzure?>(() => __config.GetObject<Pulumi.Vault.Config.Types.AuthLoginAzure>("authLoginAzure"));
        /// <summary>
        /// Login to vault using the azure method
        /// </summary>
        public static Pulumi.Vault.Config.Types.AuthLoginAzure? AuthLoginAzure
        {
            get => _authLoginAzure.Get();
            set => _authLoginAzure.Set(value);
        }

        private static readonly __Value<Pulumi.Vault.Config.Types.AuthLoginCert?> _authLoginCert = new __Value<Pulumi.Vault.Config.Types.AuthLoginCert?>(() => __config.GetObject<Pulumi.Vault.Config.Types.AuthLoginCert>("authLoginCert"));
        /// <summary>
        /// Login to vault using the cert method
        /// </summary>
        public static Pulumi.Vault.Config.Types.AuthLoginCert? AuthLoginCert
        {
            get => _authLoginCert.Get();
            set => _authLoginCert.Set(value);
        }

        private static readonly __Value<Pulumi.Vault.Config.Types.AuthLoginGcp?> _authLoginGcp = new __Value<Pulumi.Vault.Config.Types.AuthLoginGcp?>(() => __config.GetObject<Pulumi.Vault.Config.Types.AuthLoginGcp>("authLoginGcp"));
        /// <summary>
        /// Login to vault using the gcp method
        /// </summary>
        public static Pulumi.Vault.Config.Types.AuthLoginGcp? AuthLoginGcp
        {
            get => _authLoginGcp.Get();
            set => _authLoginGcp.Set(value);
        }

        private static readonly __Value<Pulumi.Vault.Config.Types.AuthLoginJwt?> _authLoginJwt = new __Value<Pulumi.Vault.Config.Types.AuthLoginJwt?>(() => __config.GetObject<Pulumi.Vault.Config.Types.AuthLoginJwt>("authLoginJwt"));
        /// <summary>
        /// Login to vault using the jwt method
        /// </summary>
        public static Pulumi.Vault.Config.Types.AuthLoginJwt? AuthLoginJwt
        {
            get => _authLoginJwt.Get();
            set => _authLoginJwt.Set(value);
        }

        private static readonly __Value<Pulumi.Vault.Config.Types.AuthLoginKerberos?> _authLoginKerberos = new __Value<Pulumi.Vault.Config.Types.AuthLoginKerberos?>(() => __config.GetObject<Pulumi.Vault.Config.Types.AuthLoginKerberos>("authLoginKerberos"));
        /// <summary>
        /// Login to vault using the kerberos method
        /// </summary>
        public static Pulumi.Vault.Config.Types.AuthLoginKerberos? AuthLoginKerberos
        {
            get => _authLoginKerberos.Get();
            set => _authLoginKerberos.Set(value);
        }

        private static readonly __Value<Pulumi.Vault.Config.Types.AuthLoginOci?> _authLoginOci = new __Value<Pulumi.Vault.Config.Types.AuthLoginOci?>(() => __config.GetObject<Pulumi.Vault.Config.Types.AuthLoginOci>("authLoginOci"));
        /// <summary>
        /// Login to vault using the OCI method
        /// </summary>
        public static Pulumi.Vault.Config.Types.AuthLoginOci? AuthLoginOci
        {
            get => _authLoginOci.Get();
            set => _authLoginOci.Set(value);
        }

        private static readonly __Value<Pulumi.Vault.Config.Types.AuthLoginOidc?> _authLoginOidc = new __Value<Pulumi.Vault.Config.Types.AuthLoginOidc?>(() => __config.GetObject<Pulumi.Vault.Config.Types.AuthLoginOidc>("authLoginOidc"));
        /// <summary>
        /// Login to vault using the oidc method
        /// </summary>
        public static Pulumi.Vault.Config.Types.AuthLoginOidc? AuthLoginOidc
        {
            get => _authLoginOidc.Get();
            set => _authLoginOidc.Set(value);
        }

        private static readonly __Value<Pulumi.Vault.Config.Types.AuthLoginRadius?> _authLoginRadius = new __Value<Pulumi.Vault.Config.Types.AuthLoginRadius?>(() => __config.GetObject<Pulumi.Vault.Config.Types.AuthLoginRadius>("authLoginRadius"));
        /// <summary>
        /// Login to vault using the radius method
        /// </summary>
        public static Pulumi.Vault.Config.Types.AuthLoginRadius? AuthLoginRadius
        {
            get => _authLoginRadius.Get();
            set => _authLoginRadius.Set(value);
        }

        private static readonly __Value<Pulumi.Vault.Config.Types.AuthLoginTokenFile?> _authLoginTokenFile = new __Value<Pulumi.Vault.Config.Types.AuthLoginTokenFile?>(() => __config.GetObject<Pulumi.Vault.Config.Types.AuthLoginTokenFile>("authLoginTokenFile"));
        /// <summary>
        /// Login to vault using
        /// </summary>
        public static Pulumi.Vault.Config.Types.AuthLoginTokenFile? AuthLoginTokenFile
        {
            get => _authLoginTokenFile.Get();
            set => _authLoginTokenFile.Set(value);
        }

        private static readonly __Value<Pulumi.Vault.Config.Types.AuthLoginUserpass?> _authLoginUserpass = new __Value<Pulumi.Vault.Config.Types.AuthLoginUserpass?>(() => __config.GetObject<Pulumi.Vault.Config.Types.AuthLoginUserpass>("authLoginUserpass"));
        /// <summary>
        /// Login to vault using the userpass method
        /// </summary>
        public static Pulumi.Vault.Config.Types.AuthLoginUserpass? AuthLoginUserpass
        {
            get => _authLoginUserpass.Get();
            set => _authLoginUserpass.Set(value);
        }

        private static readonly __Value<string?> _caCertDir = new __Value<string?>(() => __config.Get("caCertDir"));
        /// <summary>
        /// Path to directory containing CA certificate files to validate the server's certificate.
        /// </summary>
        public static string? CaCertDir
        {
            get => _caCertDir.Get();
            set => _caCertDir.Set(value);
        }

        private static readonly __Value<string?> _caCertFile = new __Value<string?>(() => __config.Get("caCertFile"));
        /// <summary>
        /// Path to a CA certificate file to validate the server's certificate.
        /// </summary>
        public static string? CaCertFile
        {
            get => _caCertFile.Get();
            set => _caCertFile.Set(value);
        }

        private static readonly __Value<Pulumi.Vault.Config.Types.ClientAuth?> _clientAuth = new __Value<Pulumi.Vault.Config.Types.ClientAuth?>(() => __config.GetObject<Pulumi.Vault.Config.Types.ClientAuth>("clientAuth"));
        /// <summary>
        /// Client authentication credentials.
        /// </summary>
        public static Pulumi.Vault.Config.Types.ClientAuth? ClientAuth
        {
            get => _clientAuth.Get();
            set => _clientAuth.Set(value);
        }

        private static readonly __Value<ImmutableArray<Pulumi.Vault.Config.Types.Headers>> _headers = new __Value<ImmutableArray<Pulumi.Vault.Config.Types.Headers>>(() => __config.GetObject<ImmutableArray<Pulumi.Vault.Config.Types.Headers>>("headers"));
        /// <summary>
        /// The headers to send with each Vault request.
        /// </summary>
        public static ImmutableArray<Pulumi.Vault.Config.Types.Headers> Headers
        {
            get => _headers.Get();
            set => _headers.Set(value);
        }

        private static readonly __Value<int?> _maxLeaseTtlSeconds = new __Value<int?>(() => __config.GetInt32("maxLeaseTtlSeconds") ?? Utilities.GetEnvInt32("TERRAFORM_VAULT_MAX_TTL") ?? 1200);
        /// <summary>
        /// Maximum TTL for secret leases requested by this provider.
        /// </summary>
        public static int? MaxLeaseTtlSeconds
        {
            get => _maxLeaseTtlSeconds.Get();
            set => _maxLeaseTtlSeconds.Set(value);
        }

        private static readonly __Value<int?> _maxRetries = new __Value<int?>(() => __config.GetInt32("maxRetries") ?? Utilities.GetEnvInt32("VAULT_MAX_RETRIES") ?? 2);
        /// <summary>
        /// Maximum number of retries when a 5xx error code is encountered.
        /// </summary>
        public static int? MaxRetries
        {
            get => _maxRetries.Get();
            set => _maxRetries.Set(value);
        }

        private static readonly __Value<int?> _maxRetriesCcc = new __Value<int?>(() => __config.GetInt32("maxRetriesCcc"));
        /// <summary>
        /// Maximum number of retries for Client Controlled Consistency related operations
        /// </summary>
        public static int? MaxRetriesCcc
        {
            get => _maxRetriesCcc.Get();
            set => _maxRetriesCcc.Set(value);
        }

        private static readonly __Value<string?> _namespace = new __Value<string?>(() => __config.Get("namespace"));
        /// <summary>
        /// The namespace to use. Available only for Vault Enterprise.
        /// </summary>
        public static string? Namespace
        {
            get => _namespace.Get();
            set => _namespace.Set(value);
        }

        private static readonly __Value<bool?> _setNamespaceFromToken = new __Value<bool?>(() => __config.GetBoolean("setNamespaceFromToken"));
        /// <summary>
        /// In the case where the Vault token is for a specific namespace and the provider namespace is not configured, use the
        /// token namespace as the root namespace for all resources.
        /// </summary>
        public static bool? SetNamespaceFromToken
        {
            get => _setNamespaceFromToken.Get();
            set => _setNamespaceFromToken.Set(value);
        }

        private static readonly __Value<bool?> _skipChildToken = new __Value<bool?>(() => __config.GetBoolean("skipChildToken"));
        /// <summary>
        /// Set this to true to prevent the creation of ephemeral child token used by this provider.
        /// </summary>
        public static bool? SkipChildToken
        {
            get => _skipChildToken.Get();
            set => _skipChildToken.Set(value);
        }

        private static readonly __Value<bool?> _skipGetVaultVersion = new __Value<bool?>(() => __config.GetBoolean("skipGetVaultVersion"));
        /// <summary>
        /// Skip the dynamic fetching of the Vault server version.
        /// </summary>
        public static bool? SkipGetVaultVersion
        {
            get => _skipGetVaultVersion.Get();
            set => _skipGetVaultVersion.Set(value);
        }

        private static readonly __Value<bool?> _skipTlsVerify = new __Value<bool?>(() => __config.GetBoolean("skipTlsVerify") ?? Utilities.GetEnvBoolean("VAULT_SKIP_VERIFY"));
        /// <summary>
        /// Set this to true only if the target Vault server is an insecure development instance.
        /// </summary>
        public static bool? SkipTlsVerify
        {
            get => _skipTlsVerify.Get();
            set => _skipTlsVerify.Set(value);
        }

        private static readonly __Value<string?> _tlsServerName = new __Value<string?>(() => __config.Get("tlsServerName"));
        /// <summary>
        /// Name to use as the SNI host when connecting via TLS.
        /// </summary>
        public static string? TlsServerName
        {
            get => _tlsServerName.Get();
            set => _tlsServerName.Set(value);
        }

        private static readonly __Value<string?> _token = new __Value<string?>(() => __config.Get("token"));
        /// <summary>
        /// Token to use to authenticate to Vault.
        /// </summary>
        public static string? Token
        {
            get => _token.Get();
            set => _token.Set(value);
        }

        private static readonly __Value<string?> _tokenName = new __Value<string?>(() => __config.Get("tokenName"));
        /// <summary>
        /// Token name to use for creating the Vault child token.
        /// </summary>
        public static string? TokenName
        {
            get => _tokenName.Get();
            set => _tokenName.Set(value);
        }

        private static readonly __Value<string?> _vaultVersionOverride = new __Value<string?>(() => __config.Get("vaultVersionOverride"));
        /// <summary>
        /// Override the Vault server version, which is normally determined dynamically from the target Vault server
        /// </summary>
        public static string? VaultVersionOverride
        {
            get => _vaultVersionOverride.Get();
            set => _vaultVersionOverride.Set(value);
        }

        public static class Types
        {

             public class AuthLogin
             {
                public string? Method { get; set; } = null!;
                public string? Namespace { get; set; } = null!;
                public ImmutableDictionary<string, string>? Parameters { get; set; } = null!;
                public string Path { get; set; }
                public bool? UseRootNamespace { get; set; }
            }

             public class AuthLoginAws
             {
                public string? AwsAccessKeyId { get; set; } = null!;
                public string? AwsIamEndpoint { get; set; } = null!;
                public string? AwsProfile { get; set; } = null!;
                public string? AwsRegion { get; set; } = null!;
                public string? AwsRoleArn { get; set; } = null!;
                public string? AwsRoleSessionName { get; set; } = null!;
                public string? AwsSecretAccessKey { get; set; } = null!;
                public string? AwsSessionToken { get; set; } = null!;
                public string? AwsSharedCredentialsFile { get; set; } = null!;
                public string? AwsStsEndpoint { get; set; } = null!;
                public string? AwsWebIdentityTokenFile { get; set; } = null!;
                public string? HeaderValue { get; set; } = null!;
                public string? Mount { get; set; } = null!;
                public string? Namespace { get; set; } = null!;
                public string Role { get; set; }
                public bool? UseRootNamespace { get; set; }
            }

             public class AuthLoginAzure
             {
                public string? ClientId { get; set; } = null!;
                public string? Jwt { get; set; } = null!;
                public string? Mount { get; set; } = null!;
                public string? Namespace { get; set; } = null!;
                public string ResourceGroupName { get; set; }
                public string Role { get; set; }
                public string? Scope { get; set; } = null!;
                public string SubscriptionId { get; set; }
                public string? TenantId { get; set; } = null!;
                public bool? UseRootNamespace { get; set; }
                public string? VmName { get; set; } = null!;
                public string? VmssName { get; set; } = null!;
            }

             public class AuthLoginCert
             {
                public string CertFile { get; set; }
                public string KeyFile { get; set; }
                public string? Mount { get; set; } = null!;
                public string? Name { get; set; } = null!;
                public string? Namespace { get; set; } = null!;
                public bool? UseRootNamespace { get; set; }
            }

             public class AuthLoginGcp
             {
                public string? Credentials { get; set; } = null!;
                public string? Jwt { get; set; } = null!;
                public string? Mount { get; set; } = null!;
                public string? Namespace { get; set; } = null!;
                public string Role { get; set; }
                public string? ServiceAccount { get; set; } = null!;
                public bool? UseRootNamespace { get; set; }
            }

             public class AuthLoginJwt
             {
                public string Jwt { get; set; }
                public string? Mount { get; set; } = null!;
                public string? Namespace { get; set; } = null!;
                public string Role { get; set; }
                public bool? UseRootNamespace { get; set; }
            }

             public class AuthLoginKerberos
             {
                public bool? DisableFastNegotiation { get; set; }
                public string? KeytabPath { get; set; } = null!;
                public string? Krb5confPath { get; set; } = null!;
                public string? Mount { get; set; } = null!;
                public string? Namespace { get; set; } = null!;
                public string? Realm { get; set; } = null!;
                public bool? RemoveInstanceName { get; set; }
                public string? Service { get; set; } = null!;
                public string? Token { get; set; } = null!;
                public bool? UseRootNamespace { get; set; }
                public string? Username { get; set; } = null!;
            }

             public class AuthLoginOci
             {
                public string AuthType { get; set; }
                public string? Mount { get; set; } = null!;
                public string? Namespace { get; set; } = null!;
                public string Role { get; set; }
                public bool? UseRootNamespace { get; set; }
            }

             public class AuthLoginOidc
             {
                public string? CallbackAddress { get; set; } = null!;
                public string? CallbackListenerAddress { get; set; } = null!;
                public string? Mount { get; set; } = null!;
                public string? Namespace { get; set; } = null!;
                public string Role { get; set; }
                public bool? UseRootNamespace { get; set; }
            }

             public class AuthLoginRadius
             {
                public string? Mount { get; set; } = null!;
                public string? Namespace { get; set; } = null!;
                public string Password { get; set; }
                public bool? UseRootNamespace { get; set; }
                public string Username { get; set; }
            }

             public class AuthLoginTokenFile
             {
                public string Filename { get; set; }
                public string? Namespace { get; set; } = null!;
                public bool? UseRootNamespace { get; set; }
            }

             public class AuthLoginUserpass
             {
                public string? Mount { get; set; } = null!;
                public string? Namespace { get; set; } = null!;
                public string? Password { get; set; } = null!;
                public string? PasswordFile { get; set; } = null!;
                public bool? UseRootNamespace { get; set; }
                public string Username { get; set; }
            }

             public class ClientAuth
             {
                public string CertFile { get; set; }
                public string KeyFile { get; set; }
            }

             public class Headers
             {
                public string Name { get; set; }
                public string Value { get; set; }
            }
        }
    }
}
