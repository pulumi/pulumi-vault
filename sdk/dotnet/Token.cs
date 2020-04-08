// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault
{
    public partial class Token : Pulumi.CustomResource
    {
        /// <summary>
        /// String containing the client token if stored in present file
        /// </summary>
        [Output("clientToken")]
        public Output<string> ClientToken { get; private set; } = null!;

        /// <summary>
        /// String containing the token display name
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// String containing the client token encrypted with the given `pgp_key` if stored in present file
        /// </summary>
        [Output("encryptedClientToken")]
        public Output<string> EncryptedClientToken { get; private set; } = null!;

        /// <summary>
        /// The explicit max TTL of this token
        /// </summary>
        [Output("explicitMaxTtl")]
        public Output<string?> ExplicitMaxTtl { get; private set; } = null!;

        /// <summary>
        /// String containing the token lease duration if present in state file
        /// </summary>
        [Output("leaseDuration")]
        public Output<int> LeaseDuration { get; private set; } = null!;

        /// <summary>
        /// String containing the token lease started time if present in state file
        /// </summary>
        [Output("leaseStarted")]
        public Output<string> LeaseStarted { get; private set; } = null!;

        /// <summary>
        /// Flag to not attach the default policy to this token
        /// </summary>
        [Output("noDefaultPolicy")]
        public Output<bool?> NoDefaultPolicy { get; private set; } = null!;

        /// <summary>
        /// Flag to create a token without parent
        /// </summary>
        [Output("noParent")]
        public Output<bool> NoParent { get; private set; } = null!;

        /// <summary>
        /// The number of allowed uses of this token
        /// </summary>
        [Output("numUses")]
        public Output<int> NumUses { get; private set; } = null!;

        /// <summary>
        /// The period of this token
        /// </summary>
        [Output("period")]
        public Output<string?> Period { get; private set; } = null!;

        /// <summary>
        /// The PGP key (base64 encoded) to encrypt the token.
        /// </summary>
        [Output("pgpKey")]
        public Output<string?> PgpKey { get; private set; } = null!;

        /// <summary>
        /// List of policies to attach to this token
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<string>> Policies { get; private set; } = null!;

        /// <summary>
        /// The renew increment
        /// </summary>
        [Output("renewIncrement")]
        public Output<int?> RenewIncrement { get; private set; } = null!;

        /// <summary>
        /// The minimal lease to renew this token
        /// </summary>
        [Output("renewMinLease")]
        public Output<int?> RenewMinLease { get; private set; } = null!;

        /// <summary>
        /// Flag to allow to renew this token
        /// </summary>
        [Output("renewable")]
        public Output<bool> Renewable { get; private set; } = null!;

        /// <summary>
        /// The token role name
        /// </summary>
        [Output("roleName")]
        public Output<string?> RoleName { get; private set; } = null!;

        /// <summary>
        /// The TTL period of this token
        /// </summary>
        [Output("ttl")]
        public Output<string?> Ttl { get; private set; } = null!;

        /// <summary>
        /// The client wrapped token.
        /// </summary>
        [Output("wrappedToken")]
        public Output<string> WrappedToken { get; private set; } = null!;

        /// <summary>
        /// The client wrapping accessor.
        /// </summary>
        [Output("wrappingAccessor")]
        public Output<string> WrappingAccessor { get; private set; } = null!;

        /// <summary>
        /// The TTL period of the wrapped token.
        /// </summary>
        [Output("wrappingTtl")]
        public Output<string?> WrappingTtl { get; private set; } = null!;


        /// <summary>
        /// Create a Token resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Token(string name, TokenArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:index/token:Token", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Token(string name, Input<string> id, TokenState? state = null, CustomResourceOptions? options = null)
            : base("vault:index/token:Token", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Token resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Token Get(string name, Input<string> id, TokenState? state = null, CustomResourceOptions? options = null)
        {
            return new Token(name, id, state, options);
        }
    }

    public sealed class TokenArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// String containing the token display name
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The explicit max TTL of this token
        /// </summary>
        [Input("explicitMaxTtl")]
        public Input<string>? ExplicitMaxTtl { get; set; }

        /// <summary>
        /// Flag to not attach the default policy to this token
        /// </summary>
        [Input("noDefaultPolicy")]
        public Input<bool>? NoDefaultPolicy { get; set; }

        /// <summary>
        /// Flag to create a token without parent
        /// </summary>
        [Input("noParent")]
        public Input<bool>? NoParent { get; set; }

        /// <summary>
        /// The number of allowed uses of this token
        /// </summary>
        [Input("numUses")]
        public Input<int>? NumUses { get; set; }

        /// <summary>
        /// The period of this token
        /// </summary>
        [Input("period")]
        public Input<string>? Period { get; set; }

        /// <summary>
        /// The PGP key (base64 encoded) to encrypt the token.
        /// </summary>
        [Input("pgpKey")]
        public Input<string>? PgpKey { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// List of policies to attach to this token
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// The renew increment
        /// </summary>
        [Input("renewIncrement")]
        public Input<int>? RenewIncrement { get; set; }

        /// <summary>
        /// The minimal lease to renew this token
        /// </summary>
        [Input("renewMinLease")]
        public Input<int>? RenewMinLease { get; set; }

        /// <summary>
        /// Flag to allow to renew this token
        /// </summary>
        [Input("renewable")]
        public Input<bool>? Renewable { get; set; }

        /// <summary>
        /// The token role name
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        /// <summary>
        /// The TTL period of this token
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        /// <summary>
        /// The TTL period of the wrapped token.
        /// </summary>
        [Input("wrappingTtl")]
        public Input<string>? WrappingTtl { get; set; }

        public TokenArgs()
        {
        }
    }

    public sealed class TokenState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// String containing the client token if stored in present file
        /// </summary>
        [Input("clientToken")]
        public Input<string>? ClientToken { get; set; }

        /// <summary>
        /// String containing the token display name
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// String containing the client token encrypted with the given `pgp_key` if stored in present file
        /// </summary>
        [Input("encryptedClientToken")]
        public Input<string>? EncryptedClientToken { get; set; }

        /// <summary>
        /// The explicit max TTL of this token
        /// </summary>
        [Input("explicitMaxTtl")]
        public Input<string>? ExplicitMaxTtl { get; set; }

        /// <summary>
        /// String containing the token lease duration if present in state file
        /// </summary>
        [Input("leaseDuration")]
        public Input<int>? LeaseDuration { get; set; }

        /// <summary>
        /// String containing the token lease started time if present in state file
        /// </summary>
        [Input("leaseStarted")]
        public Input<string>? LeaseStarted { get; set; }

        /// <summary>
        /// Flag to not attach the default policy to this token
        /// </summary>
        [Input("noDefaultPolicy")]
        public Input<bool>? NoDefaultPolicy { get; set; }

        /// <summary>
        /// Flag to create a token without parent
        /// </summary>
        [Input("noParent")]
        public Input<bool>? NoParent { get; set; }

        /// <summary>
        /// The number of allowed uses of this token
        /// </summary>
        [Input("numUses")]
        public Input<int>? NumUses { get; set; }

        /// <summary>
        /// The period of this token
        /// </summary>
        [Input("period")]
        public Input<string>? Period { get; set; }

        /// <summary>
        /// The PGP key (base64 encoded) to encrypt the token.
        /// </summary>
        [Input("pgpKey")]
        public Input<string>? PgpKey { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// List of policies to attach to this token
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// The renew increment
        /// </summary>
        [Input("renewIncrement")]
        public Input<int>? RenewIncrement { get; set; }

        /// <summary>
        /// The minimal lease to renew this token
        /// </summary>
        [Input("renewMinLease")]
        public Input<int>? RenewMinLease { get; set; }

        /// <summary>
        /// Flag to allow to renew this token
        /// </summary>
        [Input("renewable")]
        public Input<bool>? Renewable { get; set; }

        /// <summary>
        /// The token role name
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        /// <summary>
        /// The TTL period of this token
        /// </summary>
        [Input("ttl")]
        public Input<string>? Ttl { get; set; }

        /// <summary>
        /// The client wrapped token.
        /// </summary>
        [Input("wrappedToken")]
        public Input<string>? WrappedToken { get; set; }

        /// <summary>
        /// The client wrapping accessor.
        /// </summary>
        [Input("wrappingAccessor")]
        public Input<string>? WrappingAccessor { get; set; }

        /// <summary>
        /// The TTL period of the wrapped token.
        /// </summary>
        [Input("wrappingTtl")]
        public Input<string>? WrappingTtl { get; set; }

        public TokenState()
        {
        }
    }
}
