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
    /// ## Import
    /// 
    /// AWS secret backend roles can be imported using the `path`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vault:aws/secretBackendRole:SecretBackendRole role aws/roles/deploy
    /// ```
    /// </summary>
    public partial class SecretBackendRole : Pulumi.CustomResource
    {
        /// <summary>
        /// The path the AWS secret backend is mounted at,
        /// with no leading or trailing `/`s.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// Specifies the type of credential to be used when 
        /// retrieving credentials from the role. Must be one of `iam_user`, `assumed_role`, or
        /// `federation_token`.
        /// </summary>
        [Output("credentialType")]
        public Output<string> CredentialType { get; private set; } = null!;

        /// <summary>
        /// The default TTL in seconds for STS credentials. When a TTL is not specified when STS credentials are requested, and a
        /// default TTL is specified on the role, then this default TTL will be used. Valid only when credential_type is one of
        /// assumed_role or federation_token.
        /// </summary>
        [Output("defaultStsTtl")]
        public Output<int> DefaultStsTtl { get; private set; } = null!;

        /// <summary>
        /// A list of IAM group names. IAM users generated against this vault role will be added to these IAM Groups. For a
        /// credential type of assumed_role or federation_token, the policies sent to the corresponding AWS call (sts:AssumeRole or
        /// sts:GetFederation) will be the policies from each group in iam_groups combined with the policy_document and policy_arns
        /// parameters.
        /// </summary>
        [Output("iamGroups")]
        public Output<ImmutableArray<string>> IamGroups { get; private set; } = null!;

        /// <summary>
        /// The max allowed TTL in seconds for STS credentials (credentials TTL are capped to max_sts_ttl). Valid only when
        /// credential_type is one of assumed_role or federation_token.
        /// </summary>
        [Output("maxStsTtl")]
        public Output<int> MaxStsTtl { get; private set; } = null!;

        /// <summary>
        /// The name to identify this role within the backend.
        /// Must be unique within the backend.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ARN for a pre-existing policy to associate
        /// with this role. Either `policy_document` or `policy_arns` must be specified.
        /// </summary>
        [Output("policyArns")]
        public Output<ImmutableArray<string>> PolicyArns { get; private set; } = null!;

        /// <summary>
        /// The JSON-formatted policy to associate with this
        /// role. Either `policy_document` or `policy_arns` must be specified.
        /// </summary>
        [Output("policyDocument")]
        public Output<string?> PolicyDocument { get; private set; } = null!;

        /// <summary>
        /// Specifies the ARNs of the AWS roles this Vault role
        /// is allowed to assume. Required when `credential_type` is `assumed_role` and
        /// prohibited otherwise.
        /// </summary>
        [Output("roleArns")]
        public Output<ImmutableArray<string>> RoleArns { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendRole(string name, SecretBackendRoleArgs args, CustomResourceOptions? options = null)
            : base("vault:aws/secretBackendRole:SecretBackendRole", name, args ?? new SecretBackendRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendRole(string name, Input<string> id, SecretBackendRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:aws/secretBackendRole:SecretBackendRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackendRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendRole Get(string name, Input<string> id, SecretBackendRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendRole(name, id, state, options);
        }
    }

    public sealed class SecretBackendRoleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the AWS secret backend is mounted at,
        /// with no leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// Specifies the type of credential to be used when 
        /// retrieving credentials from the role. Must be one of `iam_user`, `assumed_role`, or
        /// `federation_token`.
        /// </summary>
        [Input("credentialType", required: true)]
        public Input<string> CredentialType { get; set; } = null!;

        /// <summary>
        /// The default TTL in seconds for STS credentials. When a TTL is not specified when STS credentials are requested, and a
        /// default TTL is specified on the role, then this default TTL will be used. Valid only when credential_type is one of
        /// assumed_role or federation_token.
        /// </summary>
        [Input("defaultStsTtl")]
        public Input<int>? DefaultStsTtl { get; set; }

        [Input("iamGroups")]
        private InputList<string>? _iamGroups;

        /// <summary>
        /// A list of IAM group names. IAM users generated against this vault role will be added to these IAM Groups. For a
        /// credential type of assumed_role or federation_token, the policies sent to the corresponding AWS call (sts:AssumeRole or
        /// sts:GetFederation) will be the policies from each group in iam_groups combined with the policy_document and policy_arns
        /// parameters.
        /// </summary>
        public InputList<string> IamGroups
        {
            get => _iamGroups ?? (_iamGroups = new InputList<string>());
            set => _iamGroups = value;
        }

        /// <summary>
        /// The max allowed TTL in seconds for STS credentials (credentials TTL are capped to max_sts_ttl). Valid only when
        /// credential_type is one of assumed_role or federation_token.
        /// </summary>
        [Input("maxStsTtl")]
        public Input<int>? MaxStsTtl { get; set; }

        /// <summary>
        /// The name to identify this role within the backend.
        /// Must be unique within the backend.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policyArns")]
        private InputList<string>? _policyArns;

        /// <summary>
        /// The ARN for a pre-existing policy to associate
        /// with this role. Either `policy_document` or `policy_arns` must be specified.
        /// </summary>
        public InputList<string> PolicyArns
        {
            get => _policyArns ?? (_policyArns = new InputList<string>());
            set => _policyArns = value;
        }

        /// <summary>
        /// The JSON-formatted policy to associate with this
        /// role. Either `policy_document` or `policy_arns` must be specified.
        /// </summary>
        [Input("policyDocument")]
        public Input<string>? PolicyDocument { get; set; }

        [Input("roleArns")]
        private InputList<string>? _roleArns;

        /// <summary>
        /// Specifies the ARNs of the AWS roles this Vault role
        /// is allowed to assume. Required when `credential_type` is `assumed_role` and
        /// prohibited otherwise.
        /// </summary>
        public InputList<string> RoleArns
        {
            get => _roleArns ?? (_roleArns = new InputList<string>());
            set => _roleArns = value;
        }

        public SecretBackendRoleArgs()
        {
        }
    }

    public sealed class SecretBackendRoleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the AWS secret backend is mounted at,
        /// with no leading or trailing `/`s.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// Specifies the type of credential to be used when 
        /// retrieving credentials from the role. Must be one of `iam_user`, `assumed_role`, or
        /// `federation_token`.
        /// </summary>
        [Input("credentialType")]
        public Input<string>? CredentialType { get; set; }

        /// <summary>
        /// The default TTL in seconds for STS credentials. When a TTL is not specified when STS credentials are requested, and a
        /// default TTL is specified on the role, then this default TTL will be used. Valid only when credential_type is one of
        /// assumed_role or federation_token.
        /// </summary>
        [Input("defaultStsTtl")]
        public Input<int>? DefaultStsTtl { get; set; }

        [Input("iamGroups")]
        private InputList<string>? _iamGroups;

        /// <summary>
        /// A list of IAM group names. IAM users generated against this vault role will be added to these IAM Groups. For a
        /// credential type of assumed_role or federation_token, the policies sent to the corresponding AWS call (sts:AssumeRole or
        /// sts:GetFederation) will be the policies from each group in iam_groups combined with the policy_document and policy_arns
        /// parameters.
        /// </summary>
        public InputList<string> IamGroups
        {
            get => _iamGroups ?? (_iamGroups = new InputList<string>());
            set => _iamGroups = value;
        }

        /// <summary>
        /// The max allowed TTL in seconds for STS credentials (credentials TTL are capped to max_sts_ttl). Valid only when
        /// credential_type is one of assumed_role or federation_token.
        /// </summary>
        [Input("maxStsTtl")]
        public Input<int>? MaxStsTtl { get; set; }

        /// <summary>
        /// The name to identify this role within the backend.
        /// Must be unique within the backend.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policyArns")]
        private InputList<string>? _policyArns;

        /// <summary>
        /// The ARN for a pre-existing policy to associate
        /// with this role. Either `policy_document` or `policy_arns` must be specified.
        /// </summary>
        public InputList<string> PolicyArns
        {
            get => _policyArns ?? (_policyArns = new InputList<string>());
            set => _policyArns = value;
        }

        /// <summary>
        /// The JSON-formatted policy to associate with this
        /// role. Either `policy_document` or `policy_arns` must be specified.
        /// </summary>
        [Input("policyDocument")]
        public Input<string>? PolicyDocument { get; set; }

        [Input("roleArns")]
        private InputList<string>? _roleArns;

        /// <summary>
        /// Specifies the ARNs of the AWS roles this Vault role
        /// is allowed to assume. Required when `credential_type` is `assumed_role` and
        /// prohibited otherwise.
        /// </summary>
        public InputList<string> RoleArns
        {
            get => _roleArns ?? (_roleArns = new InputList<string>());
            set => _roleArns = value;
        }

        public SecretBackendRoleState()
        {
        }
    }
}
