// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.PkiSecret
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
    ///     var intermediate = new Vault.PkiSecret.SecretBackendConfigCa("intermediate", new()
    ///     {
    ///         Backend = intermediateVaultMount.Path,
    ///         PemBundle = @"-----BEGIN RSA PRIVATE KEY-----
    /// MIIEowIBAAKCAQEAwvEHeJCXnFgi88rE1dTX6FHdBPK0wSjedh0ywVnCZxLWbBv/
    /// 5PytjTcCPdrfW7g2sfbPwOge/WF3X2KeYSP8SxZA0czmz6QDspeG921JkZWtyp5o
    /// ++N0leLTIUAhq339p3O1onAOUO1k4sHfmCwfrDpTn2hcx4URa5Pzzb1fHigusjIH
    /// 1mcGdncaA6Z2CzO1w4E8kPOUukIDrcZT4faOZrWUIQZKQw2JzTyKJ+ZMDCZq2TFz
    /// WwpL3eG48wB7J7mibFQ/9nFvxpIflBjDAZ8QiqkwYr5N0DNsTxcfTCSeubfJDCUf
    /// IWwFZhLitzwOxazazUQKXX/SPMQ1l/L9o3nnHwIDAQABAoIBAAQidJQcDPsl62fc
    /// Txxx7TpiMhvewfKu2TkMGX18V+EzxxR364+BxHSQTB3fvIkHeTGBGJrw0WdyX8PI
    /// Ja/NwZYeHLXWcLbKtcFd8WDiEoNh91Oq1HMzOc/MBcpYv94RSAX7MEkHs2YIAvHE
    /// RufFV86hVhC1d/JLYjkz5CHi+Fd9XTYjBK78tHhJd4IJPu5LYvwlmzC1zeS7s1Tg
    /// QW1FQuVDV8tWa4PMTrQHwfaGqn95AKc+tbg+ubpCiWl5bBNI3Ghuh4sAC9dMdAkd
    /// w27i29O9/Y3XJSSGUZlZqDBP4YU388RgHpzLDUxgRcaQt9vdeEz6frULPW67e9D2
    /// mPPDzjECgYEA4aPOwvnSwGoOKsS6vANGy4Ajsq09PR+1ltMJUR5kDlXGuZWI72eX
    /// 3/GAnovDuCp0tbYt0r7Fmkfel0Ore7SYM18TH5QGpPddcZLvKUf7AchCIOYY0Te3
    /// pS9+7S1lEGrLXyuox4N26Ov6wHVrmZTcQoZsDWbjYxNNsNACsiQNjGMCgYEA3SvQ
    /// Jets9e9SgNVvao2TijX+/vcNKRfcWB71T9Xc4BuSNEu5+ZLtptlwaSnVCVu1Xilk
    /// sWDh+3EhByl4EteENPvE/7A2s1sfcDOprvg0r52aBZKeTp0AukrT8+Ad4hap7g1x
    /// 2Lz11MFDkhRqt2KqQaIL+5Mq5WfptbBJ0YI7ARUCgYAD6iSfK1hlsDFYupsGwgPL
    /// agi0g97pHZC38idaOe3AdeqBs79xb9mpr/XsSj52Bn6J3IRFALxK5e5Nr4XdGo/9
    /// bCvXw2iuGgCMBOGTVMVdDY1gJr3Ne2r7Oay5Dq2PMFsg5pACDhzVA6sRBbh9LKD5
    /// on1jaiKNyHrzk1hIoOl/QwKBgA+Ov2uLbfS2yvTpDpdOMiyss603r6NOXF+Ofe8J
    /// uinBhr1K/mAB59muveuH18Z6vv1KqByaFgtb39jjH+Eja9dWRns95/sh08pOuAbo
    /// yrv3uBfgQmaBQMXZ8aLcBv4aXgWyyGlYkWpP1fL2oLMZq6RGQ9WEeqX8c0ImjmrA
    /// YGopAoGBAJZPFlZi2Rfq4MfFZp/X1/zM09hphZwkxkSI+RnsjDUjTgB8CuQul5ep
    /// KWE98yLw4C25Cqw5fKKQ2addizLnZCAIfJKVNRjYLWlWyGQydDEUzqwXlSLS9LVX
    /// LxLkWDajIyjeFn21Ttb42L9pBo3TAQIxUenom/lP2SQTvCKBiPai
    /// -----END RSA PRIVATE KEY-----
    /// -----BEGIN CERTIFICATE-----
    /// MIIDazCCAlOgAwIBAgIUahce2sCO7Bom/Rznd5HsNAlr1NgwDQYJKoZIhvcNAQEL
    /// BQAwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM
    /// GEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0xODEyMDIwMTAxNDRaFw00NjEy
    /// MTUwMTAxNDRaMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEw
    /// HwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwggEiMA0GCSqGSIb3DQEB
    /// AQUAA4IBDwAwggEKAoIBAQDC8Qd4kJecWCLzysTV1NfoUd0E8rTBKN52HTLBWcJn
    /// EtZsG//k/K2NNwI92t9buDax9s/A6B79YXdfYp5hI/xLFkDRzObPpAOyl4b3bUmR
    /// la3Knmj743SV4tMhQCGrff2nc7WicA5Q7WTiwd+YLB+sOlOfaFzHhRFrk/PNvV8e
    /// KC6yMgfWZwZ2dxoDpnYLM7XDgTyQ85S6QgOtxlPh9o5mtZQhBkpDDYnNPIon5kwM
    /// JmrZMXNbCkvd4bjzAHsnuaJsVD/2cW/Gkh+UGMMBnxCKqTBivk3QM2xPFx9MJJ65
    /// t8kMJR8hbAVmEuK3PA7FrNrNRApdf9I8xDWX8v2jeecfAgMBAAGjUzBRMB0GA1Ud
    /// DgQWBBQXGfrns8OqxTGKsXG5pDZS/WyyYDAfBgNVHSMEGDAWgBQXGfrns8OqxTGK
    /// sXG5pDZS/WyyYDAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQCt
    /// 8aUX26cl2PgdIEByZSHAX5G+2b0IEtTclPkl4uDyyKRY4dVq6gK3ueVSU5eUmBip
    /// JbV5aRetovGOcV//8vbxkZm/ntQ8Oo+2sfGR5lIzd0UdlOr5pkD6g3bFy/zJ+4DR
    /// DAe8fklUacfz6CFmD+H8GyHm+fKmF+mjr4oOGQW6OegRDJHuiipUk2lJyuXdlPSa
    /// FpNRO2sGbjn000ANinFgnFiVzGDnx0/G1Kii/6GWrI6rrdVmXioQzF+8AloWckeB
    /// +hbmbwkwQa/JrLb5SWcBDOXSgtn1Li3XF5AQQBBjA3pOlyBXqnI94Irw89Lv9uPT
    /// MUR4qFxeUOW/GJGccMUd
    /// -----END CERTIFICATE-----
    /// ",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn =
    ///         {
    ///             intermediateVaultMount,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [VaultResourceType("vault:pkiSecret/secretBackendConfigCa:SecretBackendConfigCa")]
    public partial class SecretBackendConfigCa : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The PKI secret backend the resource belongs to.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// The key and certificate PEM bundle
        /// </summary>
        [Output("pemBundle")]
        public Output<string> PemBundle { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendConfigCa resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendConfigCa(string name, SecretBackendConfigCaArgs args, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendConfigCa:SecretBackendConfigCa", name, args ?? new SecretBackendConfigCaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendConfigCa(string name, Input<string> id, SecretBackendConfigCaState? state = null, CustomResourceOptions? options = null)
            : base("vault:pkiSecret/secretBackendConfigCa:SecretBackendConfigCa", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "pemBundle",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecretBackendConfigCa resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendConfigCa Get(string name, Input<string> id, SecretBackendConfigCaState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendConfigCa(name, id, state, options);
        }
    }

    public sealed class SecretBackendConfigCaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The PKI secret backend the resource belongs to.
        /// </summary>
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

        [Input("pemBundle", required: true)]
        private Input<string>? _pemBundle;

        /// <summary>
        /// The key and certificate PEM bundle
        /// </summary>
        public Input<string>? PemBundle
        {
            get => _pemBundle;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _pemBundle = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public SecretBackendConfigCaArgs()
        {
        }
        public static new SecretBackendConfigCaArgs Empty => new SecretBackendConfigCaArgs();
    }

    public sealed class SecretBackendConfigCaState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The PKI secret backend the resource belongs to.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("pemBundle")]
        private Input<string>? _pemBundle;

        /// <summary>
        /// The key and certificate PEM bundle
        /// </summary>
        public Input<string>? PemBundle
        {
            get => _pemBundle;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _pemBundle = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public SecretBackendConfigCaState()
        {
        }
        public static new SecretBackendConfigCaState Empty => new SecretBackendConfigCaState();
    }
}
